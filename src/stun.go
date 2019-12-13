package main

import (
	"gortc.io/stun"
	"net"
)

func Stun(mode Mode, server string) (net.IP, error) {
	var network string
	switch mode {
	case MODE_AUTO:
		network = "udp"
	case MODE_IPv4:
		network = "udp4"
	case MODE_IPv6:
		network = "udp6"
	}

	if server == "-" {
		server = "stun.l.google.com:19302"
	}

	// Creating a "connection" to STUN server.
	c, err := stun.Dial(network, server)
	if err != nil {
		return nil, err
	}
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	retChan := make(chan stun.Event)

	// Sending request to STUN server, waiting for response message.
	if err := c.Start(message, func(res stun.Event) {
		retChan <- res
	}); err != nil {
		return nil, err
	}

	res := <-retChan

	if res.Error != nil {
		return nil, res.Error
	}

	// Decoding XOR-MAPPED-ADDRESS attribute from message.
	var xorAddr stun.XORMappedAddress
	if err := xorAddr.GetFrom(res.Message); err != nil {
		return nil, err
	}
	return xorAddr.IP, nil
}
