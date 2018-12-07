package main

import (
	"context"
	"errors"
	"math/rand"
	"net"
)

func dialContextIPv4(ctx context.Context, network, address string) (net.Conn, error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
		network = "tcp4"
	case "udp", "udp4", "udp6":
		network = "udp4"
	}

	return (&net.Dialer{}).DialContext(ctx, network, address)
}

func dialContextIPv6(ctx context.Context, network, address string) (net.Conn, error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
		network = "tcp6"
	case "udp", "udp4", "udp6":
		network = "udp6"
	}
	return (&net.Dialer{}).DialContext(ctx, network, address)
}

func dialContextDualStack(ctx context.Context, network, address string) (net.Conn, error) {
	return (&net.Dialer{DualStack: true}).DialContext(ctx, network, address)
}

func OpenDnsDnsQuery(mode Mode, server string) (net.IP, error) {
	ctx := context.Background()

	defaultAddresses := []string{
		"resolver1.opendns.com:53",
		"resolver2.opendns.com:53",
	}
	n := rand.Int() % len(defaultAddresses)

	apiEndpoint := server
	if apiEndpoint == "-" {
		apiEndpoint = defaultAddresses[n]
	}

	resolver := &net.Resolver{
		PreferGo: true,
	}
	switch mode {
	case MODE_IPv4:
		resolver.Dial = dialContextIPv4
	case MODE_IPv6:
		resolver.Dial = dialContextIPv6
	default:
		resolver.Dial = dialContextDualStack
	}

	ips, err := resolver.LookupIPAddr(ctx, "myip.opendns.com")
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, errors.New("server returned no IP address")
	}
	return ips[0].IP, nil
}
