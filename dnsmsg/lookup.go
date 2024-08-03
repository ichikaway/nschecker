package dnsmsg

import (
	"errors"
	"net"
	"nschecker/printer"
	"strings"
)

func Lookup(domainName string) ([]string, error) {
	servers, err := getAuthorityServerName(domainName)
	if err != nil {
		printer.Printf(" ..lookup from local DNS cache server.\n\n")
		return lookupFromDnsCacheServer(domainName)
	}

	var ret []string
	var cnt int = 0
	for index, server := range servers {
		printer.Printf(" ..lookup from DNS root server: %s \n\n", server)
		ret, err = lookupFromDnsRoot(domainName, server)
		if err != nil {
			printer.Printf("  [%d] error(lookup from root server): %s \n", index, err)
		} else {
			break
		}
		cnt++
		if cnt > 3 {
			printer.Printf("  lookup exceeded(lookup from root server) \n")
			return ret, err
		}
	}
	return ret, err
}

// get DNS server for looking up name from DNS root servers.
func getAuthorityServerName(domainName string) ([]string, error) {
	labels := strings.Split(domainName, ".")

	if len(labels) == 1 {
		return []string{}, errors.New("not support domain level.")
	}
	gTldLabel := labels[len(labels)-1] //get last value in array
	servers, err := getTldNameServer(gTldLabel)
	if err != nil {
		return []string{}, errors.New("not support TLD name.")
	}
	return servers, nil
}

// get TLD NS server name from full resolver server.
func getTldNameServer(tldName string) ([]string, error) {
	servers, err := lookupFromDnsCacheServer(tldName)
	if err != nil {
		return []string{}, err
	}
	return servers, nil
}

func lookupFromDnsCacheServer(domainName string) ([]string, error) {
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

func lookupFromDnsRoot(domainName string, dnsServer string) ([]string, error) {
	id := NewId()
	h := NewHeader(id)
	header := NewHeaderPayload(h)
	q := NewQuestion(domainName, TypeNS)
	question := NewQuestionPayload(q)
	req := append(header, question...)

	// sending UDP
	message, err := Send(dnsServer+":53", req)
	if err != nil {
		return nil, err
	}

	if checkValidDnsId(message, id) == false {
		return nil, errors.New("Not match DNS ID value.")
	}

	ret, err := getNsListFromDnsResponse(message)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
