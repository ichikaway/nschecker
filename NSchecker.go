package main

import (
	"flag"
	"nschecker/checker"
	"nschecker/notification"
	"nschecker/printer"
	"os"
)

var VERSION = "1.0.2"

func showError() {
	printer.ErrorPrintf("USAGE: go run NsCheck.go -type NS -domain domainName -expect 'ns records' \n")
	printer.ErrorPrintf(" or  \n")
	printer.ErrorPrintf("USAGE (Deplicated): go run NsCheck.go type(NS or MX) 'domain' 'ns records' \n")
	os.Exit(1)
}

func main() {
	var qType string
	var domainName string
	var nsListString string

	if len(os.Args) < 4 {
		showVersion()
		showError()
	}

	if len(os.Args) != 4 {
		qType2 := flag.String("type", "NS", "type: NS or MX")
		domainName2 := flag.String("domain", "", "domain name: vaddy.net")
		nsListString2 := flag.String("expect", "", "ex: 'ns1.vaddy.net, ns2.vaddy.net'")
		mode := flag.String("mode", "", "optional: silent")
		flag.Parse()

		qType = *qType2
		domainName = *domainName2
		nsListString = *nsListString2

		if *mode == "silent" {
			printer.Printf = printer.PrinterNothing
		}
	} else {
		qType = os.Args[1]
		domainName = os.Args[2]
		nsListString = os.Args[3]
	}

	showVersion()

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
	printer.Printf(" - Domain: %s\n", domainName)
	printer.Printf(" - Type: %s\n", qType)
	printer.Printf(" - List: %s\n", nsListString)
	printer.Printf("")
}

func showVersion() {
	printer.Printf("=== NSchecker Version: %s === \n", VERSION)
}
