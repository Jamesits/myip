package main

import (
	"context"
	"errors"
	"math/rand"
	"net"
)

func dialContextFactory(mode Mode, server string) func(context.Context, string, string) (net.Conn, error) {
	defaultAddresses := []string{
		"resolver1.opendns.com:53",
		"resolver2.opendns.com:53",
	}
	n := rand.Intn(len(defaultAddresses))

	apiEndpoint := server
	if apiEndpoint == "-" {
		apiEndpoint = defaultAddresses[n]
	}

	var ret func(context.Context, string, string) (net.Conn, error)
	switch mode {
	case MODE_IPv4:
		ret = func(ctx context.Context, network, address string) (net.Conn, error) {
			switch network {
			case "tcp", "tcp4", "tcp6":
				network = "tcp4"
			case "udp", "udp4", "udp6":
				network = "udp4"
			}

			return (&net.Dialer{}).DialContext(ctx, network, apiEndpoint)
		}
	case MODE_IPv6:
		ret = func(ctx context.Context, network, address string) (net.Conn, error) {
			switch network {
			case "tcp", "tcp4", "tcp6":
				network = "tcp6"
			case "udp", "udp4", "udp6":
				network = "udp6"
			}
			return (&net.Dialer{}).DialContext(ctx, network, apiEndpoint)
		}
	default:
		ret = func(ctx context.Context, network, address string) (net.Conn, error) {
			return (&net.Dialer{DualStack: true}).DialContext(ctx, network, apiEndpoint)
		}
	}

	return ret
}

func OpenDnsDnsQuery(mode Mode, server string) (net.IP, error) {
	ctx := context.Background()

	resolver := &net.Resolver{
		PreferGo: true,
	}

	resolver.Dial = dialContextFactory(mode, server)

	ips, err := resolver.LookupIPAddr(ctx, "myip.opendns.com")
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, errors.New("server returned no IP address")
	}
	return ips[0].IP, nil
}
