package main

import (
	"context"
	"math/rand"
	"net"
)

// override the network and address of the DNS request
func CustomDNSDialerFactory(address string) func(context.Context, string, string) (net.Conn, error) {
	return func (ctx context.Context, network string, address string) (net.Conn, error) {
		d := net.Dialer{}
		return d.DialContext(ctx, network, address)
		}
}

func getOpenDnsResolverIpv4() string {
	addresses := []string{
		"208.67.222.222:53",
		"208.67.220.220:53",
	}
	n := rand.Int() % len(addresses)
	return addresses[n]
}

func getOpenDnsResolverIpv6() string {
	addresses := []string{
		"[2620:119:35::35]:53",
		"[2620:119:53::53]:53",
	}
	n := rand.Int() % len(addresses)
	return addresses[n]
}

func OpenDnsDnsQuery(mode int, server string) (net.IP, error) {
	var apiEndpoint string
	if server == "-" {
		switch mode {
		case MODE_AUTO:
			apiEndpoint = "resolver1.opendns.com:53"
		case MODE_IPv4:
			apiEndpoint = getOpenDnsResolverIpv4()
		case MODE_IPv6:
			apiEndpoint = getOpenDnsResolverIpv6()
		}
	} else {
		apiEndpoint = server
	}

	resolver := net.Resolver {
		PreferGo:true,
		Dial:CustomDNSDialerFactory(apiEndpoint),
	}
	ctx := context.Background()

	ips, err := resolver.LookupIPAddr(ctx, "myip.opendns.com")
	if err != nil {
		return nil, err
	}
	return ips[0].IP, nil
}