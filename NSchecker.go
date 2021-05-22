package main

import (
	"flag"
	"fmt"
	"nschecker/checker"
	"nschecker/notification"
	"os"
)

var VERSION = "1.0.2"

func showError() {
	fmt.Printf("USAGE: go run NsCheck.go -type NS -domain domainName -expect 'ns records' \n")
	fmt.Printf(" or  \n")
	fmt.Printf("USAGE (Deplicated): go run NsCheck.go type(NS or MX) 'domain' 'ns records' \n")
	os.Exit(1)
}

func main() {
	var qType string
	var domainName string
	var nsListString string

	fmt.Printf("=== NSchecker Version: %s === \n", VERSION)
	if len(os.Args) < 4 {
		showError()
	}

	if len(os.Args) != 4 {
		qType2 := flag.String("type", "NS", "type: NS or MX")
		domainName2 := flag.String("domain", "", "domain name: vaddy.net")
		nsListString2 := flag.String("expect", "", "ex: 'ns1.vaddy.net, ns2.vaddy.net'")
		flag.Parse()

		qType = *qType2
		domainName = *domainName2
		nsListString = *nsListString2
	} else {
		qType = os.Args[1]
		domainName = os.Args[2]
		nsListString = os.Args[3]
	}

	if qType != "NS" && qType != "MX" {
		showError()
	}

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
