package checker

import (
	"errors"
	"fmt"
	"net"
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

func getNsRecords(domainName string) ([]string, error) {
	var ret []string
	nss, err := net.LookupNS(domainName)
	if err != nil {
		return nil, errors.New("NS Lookup Error.\n")
	}
	for _, ns := range nss {
		ret = append(ret, ns.Host)
	}
	return ret, nil
}

func getMxRecords(domainName string) ([]string, error) {
	var ret []string
	nss, err := net.LookupMX(domainName)
	if err != nil {
		return nil, errors.New("MX Lookup Error.\n")
	}
	for _, ns := range nss {
		ret = append(ret, ns.Host)
	}
	return ret, nil
}

func CheckRecord(qType string, domainName string, expectString string) error {
	var records []string
	var err error

	if qType == "NS" {
		records, err = getNsRecords(domainName)
	}
	if qType == "MX" {
		records, err = getMxRecords(domainName)
	}

	nsValueList := strings.Split(expectString, ",")

	if err != nil {
		return errors.New("Error: Lookup Error.\n")
	}
	if len(records) == 0 {
		return errors.New("Error: no NS record.\n")
	}

	for _, record := range records {
		if !in_array(record, nsValueList) {
			fmt.Printf("Error actual DNS records:\n")
			for _, actualRecord := range records {
				fmt.Printf("  %s\n", actualRecord)
			}
			return errors.New("Error: not match DNS record value.\n")
		}
	}
	return nil
}
