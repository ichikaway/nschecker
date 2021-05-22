package checker

import (
	"errors"
	"fmt"
	"net"
	"nschecker/dnsmsg"
	"strings"
)

func in_array(str string, list []string) bool {
	for _, v := range list {
		v = strings.TrimSpace(v)
		if strings.ToLower(v) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

func getExpectRecordValueList(expectString string) []string {
	list := strings.Split(expectString, ",")
	//Add last period for each record value
	for i, v := range list {
		v = strings.TrimSpace(v)
		v = strings.TrimRight(v, ".")
		v = v + "."
		list[i] = v
	}
	return list
}

// for getting NS records, try to get NS rr from DNS root servers.
func getNsRecords(domainName string) ([]string, error) {
	return dnsmsg.Lookup(domainName)
}

func getMxRecords(domainName string) ([]string, error) {
	var ret []string
	fmt.Print(" ..lookup from local DNS cache server.\n\n")
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

	nsValueList := getExpectRecordValueList(expectString)

	if err != nil {
		return errors.New("Error: Lookup Error.\n")
	}
	if len(records) == 0 {
		return errors.New("Error: no record.\n")
	}

	for _, record := range records {
		if !in_array(record, nsValueList) {
			message := "Error actual DNS records(" + qType + "):\n"
			for _, actualRecord := range records {
				message += fmt.Sprintf("  %s\n", actualRecord)
			}

			message += "\nExpect DNS records:\n"
			for _, expectRecord := range nsValueList {
				message += fmt.Sprintf("  %s\n", strings.TrimSpace(expectRecord))
			}

			return errors.New("Error: not match DNS record value.\n" + message)
		}
	}
	fmt.Println("PASS.  No problems.")
	return nil
}
