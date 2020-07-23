package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

var VERSION = "0.01"

func in_array(str string, list []string) bool {
	for _, v := range list {
		v = strings.TrimSpace(v)
		if v == str {
			return true
		}
	}
	return false
}

func checkNs(domainName string, expectString string) error {
	nsValueList := strings.Split(expectString, ",")

	nss, err := net.LookupNS(domainName)
	if err != nil {
		return errors.New("Lookup Error.\n")
	}
	if len(nss) == 0 {
		return errors.New("no NS record.\n")
	}
	if len(nsValueList) != len(nss) {
		return errors.New("Warging. no Match Record count.\n")
	}

	for _, ns := range nss {
		if !in_array(ns.Host, nsValueList) {
			return errors.New("Warging. no Match Record value.\n")
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("NsCheck Version: %s\n", VERSION)
		fmt.Printf("USAGE: go run NsCheck.go 'domain' 'ns records' \n")
		os.Exit(1)
	}

	var domainName string = os.Args[1]
	var nsListString string = os.Args[2]

	err := checkNs(domainName, nsListString)
	if err != nil {
		panic(err)
	}
}
