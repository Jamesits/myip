package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func OpenDnsHttpApi(mode Mode, server string) (net.IP, error) {
	apiEndpoint := "https://diagnostic.opendns.com/myip"

	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP Error %d", resp.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ip := net.ParseIP(FilterIP(string(bodyBytes)))
	return ip, nil
}
