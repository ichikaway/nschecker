package dnsmsg

import (
	"testing"
)

type A struct {
	Int32Field int32
	ByteField  byte
}

/*
func TestConnect(t *testing.T) {
	//Send([]byte("Hello from Client"))
	//Send([]byte{0xe3, 0x81, 0x82})
	//a := A{Int32Field: 0x123456, ByteField: 0xFF}
	//a := A{Int32Field: 0xe38182, ByteField: 0xFF}

	//a := try("www.vaddy.net")
	a := try("a.net")
	fmt.Printf("%v\n",a)
	b := bytes.NewBuffer([]byte(""))

	// 構造体をバイナリにする
	binary.Write(b, binary.BigEndian, &a)
	//binary.Write(b, binary.LittleEndian, &a)
	fmt.Printf("%x\n",b.Bytes())
	Send(b.Bytes())
}
*/
func TestConnect(t *testing.T) {
	id := uint16(60000)
	h := NewHeader(id)
	header := NewHeaderPayload(h)
	q := NewQuestion("bitforest.jp", TypeNS)
	//q := NewQuestion("vaddy.net", TypeNS)
	question := NewQuestionPayload(q)
	req := append(header, question...)
	//Send("192.168.12.1:53", req) //local dns cache
	Send("203.119.1.1:53", req) //a.dns.jp
}
