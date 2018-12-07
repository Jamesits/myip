package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func IpSb(mode int, server string) (net.IP, error) {
	var apiEndpoint string
	if server == "-" {
		switch mode {
		case MODE_AUTO:
			apiEndpoint = "https://api.ip.sb/ip"
		case MODE_IPv4:
			apiEndpoint = "https://api-ipv4.ip.sb/ip"
		case MODE_IPv6:
			apiEndpoint = "https://api-ipv6.ip.sb/ip"
		}
	} else {
		apiEndpoint = server
	}

	var client http.Client
	resp, err := client.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		ip := net.ParseIP(FilterIP(string(bodyBytes)))
		return ip, nil
	}

	return nil, errors.New(fmt.Sprintf("HTTP Error %d", resp.StatusCode))
}
