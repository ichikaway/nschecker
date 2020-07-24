package main

import (
	"fmt"
	"nscheck/checker"
	"os"
)

var VERSION = "0.01"

func showError() {
	fmt.Printf("NsCheck Version: %s\n", VERSION)
	fmt.Printf("USAGE: go run NsCheck.go Type(NS or MX) 'domain' 'ns records' \n")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 4 {
		showError()
	}

	var qType string = os.Args[1]
	if qType != "NS" && qType != "MX" {
		showError()
	}

	var domainName string = os.Args[2]
	var nsListString string = os.Args[3]

	err := checker.CheckRecord(qType, domainName, nsListString)
	if err != nil {
		panic(err)
	}
}
