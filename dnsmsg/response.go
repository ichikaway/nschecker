package dnsmsg

import (
	"encoding/binary"
	"errors"
	"golang.org/x/net/dns/dnsmessage"
)

func getNsListFromDnsResponse(message []byte) ([]string, error) {
	//fmt.Printf("dnsresponse: %s\n", message)
	//fmt.Printf("dnsresponse: %#08b\n", message)
	//fmt.Printf("dnsresponse: %#b\n", message)
	//fmt.Printf("dnsresponse: %08b\n", message)
	//fmt.Printf("dnsresponse: %x\n", message)
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

// こちらで生成したDNS IDとレスポンスでセットされたDNS IDが一致するかチェックする
// ここが一致しないと不正なDNSレスポンスを受け取っている可能性があるため
func checkValidDnsId(message []byte, id uint16) bool {
	receiveId := message[:2]
	result := binary.BigEndian.Uint16(receiveId)
	if result == id {
		return true
	}
	//fmt.Println(result)
	return false
}
