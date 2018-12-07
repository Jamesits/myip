package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	ipv4 := flag.Bool("4", false, "Prefer IPv4")
	ipv6 := flag.Bool("6", false, "Prefer IPv6")
	method := flag.String("method", "stun", "Method [STUN|OpenDNS|OpenDNS-API|ip.sb]")
	server := flag.String("server", "-", "Server URI, if applicable")
	flag.Parse()

	parsedMethod := strings.ToLower(*method)

	// select IP mode
	mode := MODE_AUTO
	if *ipv4 == true && *ipv6 == false {
		mode = MODE_IPv4
	} else if *ipv4 == false && *ipv6 == true {
		mode = MODE_IPv6
	}

	var ret net.IP
	var err error

	switch parsedMethod {
	case "stun":
		ret, err = Stun(mode, *server)
	case "ip.sb":
		ret, err = IpSb(mode, *server)
	case "opendns-api":
		ret, err = OpenDnsHttpApi(mode, *server)
	case "opendns":
		ret, err = OpenDnsDnsQuery(mode, *server)
	default:
		panic("Unknown method")
	}

	if err != nil {
		panic(err)
	}

	if ret == nil {
		panic(errors.New("Failed getting IP address"))
	}

	fmt.Println(ret)
}
