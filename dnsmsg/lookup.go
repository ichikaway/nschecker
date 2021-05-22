package dnsmsg

import (
	"errors"
	"net"
	"nschecker/printer"
	"strings"
)

func Lookup(domainName string) ([]string, error) {
	server, err := getAuthorityServerName(domainName)
	if err != nil {
		printer.Printf(" ..lookup from local DNS cache server.\n\n")
		return lookupFromDnsCacheServer(domainName)
	}
	printer.Printf(" ..lookup from DNS root server: %s \n\n", server)
	return lookupFromDnsRoot(domainName, server)
}

// get DNS server for looking up name from DNS root servers.
func getAuthorityServerName(domainName string) (string, error) {
	labels := strings.Split(domainName, ".")

	if len(labels) == 1 {
		return "", errors.New("not support domain level.")
	}
	gTldLabel := labels[len(labels)-1] //get last value in array
	server, err := getTldNameServer(gTldLabel)
	if err != nil {
		return "", errors.New("not support TLD name.")
	}
	return server, nil
}

// get TLD NS server name from full resolver server.
func getTldNameServer(tldName string) (string, error) {
	servers, err := lookupFromDnsCacheServer(tldName)
	if err != nil {
		return "", err
	}
	return servers[0], nil
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
