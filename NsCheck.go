package main

import (
	"net"
	"os"
	"strings"
)

func in_array(str string, list []string) bool {
	for _, v := range list {
		v = strings.TrimSpace(v)
		if v == str {
			return true
		}
	}
	return false
}

func checkNs(domainName string, expectString string) (string, bool) {
	nsValueList := strings.Split(expectString, ",")

	nss, err := net.LookupNS(domainName)
	if err != nil {
		return "Lookup Error.\n", false
	}
	if len(nss) == 0 {
		return "no NS record.\n", false
	}
	if len(nsValueList) != len(nss) {
		return "Warging. no Match Record count.\n", false
	}

	for _, ns := range nss {
		if !in_array(ns.Host, nsValueList) {
			return "Warging. no Match Record value.\n", false
		}
	}
	return "ok\n", true
}

func main() {
	var domainName string = os.Args[1]
	var nsListString string = os.Args[2]

	result, err := checkNs(domainName, nsListString)
	if err == false {
		panic(result)
	}
}
