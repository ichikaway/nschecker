package dnsmsg

import (
	"errors"
	"golang.org/x/net/dns/dnsmessage"
)

func getNsListFromDnsResponse(message []byte) ([]string, error) {
	var ret []string
	var m dnsmessage.Message
	err := m.Unpack(message)
	if err != nil {
		return nil, err
	}

	if len(m.Authorities) == 0 {
		return nil, errors.New("NS Lookup Error. No NS servers from DNS root.\n")
	}

	for _, authrotiy := range m.Authorities {
		rr := authrotiy.Body.(*dnsmessage.NSResource)
		ret = append(ret, rr.NS.String())
	}
	return ret, nil
}
