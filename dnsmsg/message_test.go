package dnsmsg

import (
	"bytes"
	"fmt"
	"testing"
)

/*
func TestMessage(t *testing.T) {
	q := try("www.vaddy.net")
	fmt.Printf("%v", q)
	id, udpReq, _, _ := newRequest(q)
	fmt.Printf("%d", id)
	fmt.Printf("%v", udpReq)
}
*/
/*
func TestMessage(t *testing.T) {
	id := uint16(60000)
	h := NewHeader(id)
	req := NewHeaderPayload(h)
	//fmt.Printf("%v\n", h)

	// 構造体をバイナリにする
	//b := bytes.NewBuffer([]byte(""))
	//binary.Write(b, binary.BigEndian, h)
	//binary.Write(b, binary.LittleEndian, &a)

	//fmt.Printf("hex: %x\n",b.Bytes())
	//fmt.Printf("hex: %x\n",b)
	//fmt.Printf("bin: %b\n",b.Bytes())
	//fmt.Printf("len: %d\n",len(b.Bytes()))

	fmt.Printf("bin: %b\n", req)
	// ea60 0100 0001 0000 0000 0000

	//hoge := []byte{0xea, 0x60}
	//fmt.Printf("hoge: %d\n", binary.BigEndian.Uint16(hoge))

	q := NewQuestion("vaddy.net", TypeNS)
	req2 := NewQuestionPayload(q)
	fmt.Printf("bin question: %b\n", req2)
}
*/

func TestCreateDomainLabel2ndLevel(t *testing.T) {
	var name string = "vaddy.net"
	var expect []byte = []byte{0x05, 0x76, 0x61, 0x64, 0x64, 0x79, 0x03, 0x6e, 0x65, 0x74} //5vaddy3net
	result := createDomainLabel(name)
	if bytes.Compare(result, expect) != 0 {
		fmt.Printf("expect: %b\n", expect)
		fmt.Printf("result: %b\n", result)
		fmt.Printf("result: %x\n", result)
		t.Fail()
	}
}
func TestCreateDomainLabel3rdLevel(t *testing.T) {
	var name string = "aa.vaddy.net"
	var expect []byte = []byte{0x02, 0x61, 0x61, 0x05, 0x76, 0x61, 0x64, 0x64, 0x79, 0x03, 0x6e, 0x65, 0x74} //2aa5vaddy3net
	result := createDomainLabel(name)
	if bytes.Compare(result, expect) != 0 {
		fmt.Printf("expect: %b\n", expect)
		fmt.Printf("result: %b\n", result)
		fmt.Printf("result: %x\n", result)
		t.Fail()
	}
}
