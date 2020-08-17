package main

import (
	"fmt"
	"nscheck/checker"
	"nscheck/notification"
	"os"
)

var VERSION = "0.03"

func showError() {
	fmt.Printf("USAGE: go run NsCheck.go Type(NS or MX) 'domain' 'ns records' \n")
	os.Exit(1)
}

func main() {
	fmt.Printf("=== NSchecker Version: %s === \n", VERSION)
	if len(os.Args) < 4 {
		showError()
	}

	var qType string = os.Args[1]
	if qType != "NS" && qType != "MX" {
		showError()
	}

	var domainName string = os.Args[2]
	var nsListString string = os.Args[3]

	infoDump(domainName, qType, nsListString)

	err := checker.CheckRecord(qType, domainName, nsListString)
	if err != nil {
		notification.PostSlack("NSchecker (Ver. "+VERSION+")", err.Error(), domainName, qType)
		panic(err)
	}
}

func infoDump(domainName string, qType string, nsListString string) {
	fmt.Println(" - Domain: " + domainName)
	fmt.Println(" - Type: " + qType)
	fmt.Println(" - List: " + nsListString)
	fmt.Println("")
}
