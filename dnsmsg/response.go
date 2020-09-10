package dnsmsg

import (
	"encoding/binary"
	"errors"
	"fmt"
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

func getNsListFromDnsResponse2(message []byte) ([]string, error) {
	const headerByteLen int = 12
	const typeByteLen int = 2
	const classByteLen int = 2
	const nullByteLen int = 1

	var ret []string
	var readCounter int = 0 //現在メッセージの何バイト目まで読み込んだかを示すカウンター

	header := message[:headerByteLen]
	readCounter = headerByteLen

	var qCount uint16 = binary.BigEndian.Uint16(header[4:6])
	var answerCount uint16 = binary.BigEndian.Uint16(header[6:8])
	var nsCount uint16 = binary.BigEndian.Uint16(header[8:10])

	if qCount != 1 {
		return nil, errors.New("Question count needs 1")
	}
	if answerCount != 0 {
		return nil, errors.New("Answer count needs 0")
	}
	if nsCount == 0 {
		return nil, errors.New("Authrity count needs more than 1")
	}

	// ---- read Question section ----
	qName, qRead, err := readName(message, readCounter)
	if err != nil {
		return nil, err
	}
	readCounter = readCounter + qRead + nullByteLen + typeByteLen + classByteLen
	fmt.Printf("%s, %d, %d", qName, qRead, readCounter)

	// ----- read Authority section ----
	// todo nsCountの数だけreadNameしてretの配列に入れる
	// nsのreadNameは名前圧縮があるのでreadNameを再帰的な動きにさせる

	return ret, nil
}

// read domain name
// name ex. vaddy.net
// readByte means read byte size
func readName(message []byte, readCounter int) (name string, readByte int, err error) {
	const dot byte = 0x2e
	data := message[readCounter:]
	var labelCount uint8 = 0
	var nameByte []byte = make([]byte, 0, 50)
	for readByte, byteData := range data {
		if byteData == 0x00 {
			name = string(nameByte[1:]) //先頭のドットは不要
			return name, readByte, nil
		}
		if labelCount == 0 {
			//labelCount=0はラベル文字数を読み取り、ドットの文字を連結
			nameByte = append(nameByte, dot)
			labelCount = byteData
		} else {
			nameByte = append(nameByte, byteData)
			labelCount--
		}
	}
	return "", 0, errors.New("No terminate byte.")
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
