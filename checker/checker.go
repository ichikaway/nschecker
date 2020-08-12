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

	/*
		r := &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Millisecond * time.Duration(10000),
				}
				return d.DialContext(ctx, "udp", "a.gtld-servers.net:53")
				//return d.DialContext(ctx, "udp", "ns-700.awsdns-23.net:53")
				//return d.DialContext(ctx, "udp", "192.55.83.30:53")
			},
		}
		nss, err := r.LookupNS(context.Background(), domainName)
	*/

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
	return nil
}
