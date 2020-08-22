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
func TestCreateHeaderPayload(t *testing.T) {
	var id uint16 = 65535
	h := NewHeader(id)
	result := NewHeaderPayload(h)

	//DNS header ID=65535, RD=0, QD=1
	var expect []byte = []byte{
		0b11111111, 0b11111111, //ID 16bit 65535
		0b00000000,             //h_16 8bit
		0b00000000,             //h_24 8bit
		0b00000000, 0b00000001, //QD Question 16bit
		0b00000000, 0b00000000, //AN Answer 16bit
		0b00000000, 0b00000000, //NS Authority 16bit
		0b00000000, 0b00000000, //AR additional 16bit
	}
	if bytes.Compare(result, expect) != 0 {
		fmt.Printf("expect: %b\n", expect)
		fmt.Printf("result: %b\n", result)
		fmt.Printf("result: %x\n", result)
		t.Fail()
	}
}

func TestCreateQuestionPayload(t *testing.T) {
	q := NewQuestion("vaddy.net", TypeNS)
	result := NewQuestionPayload(q)

	//5vaddy3net0000(type 02)(class 01)  type 02=NS, class 01=inet
	var expect []byte = []byte{0x05, 0x76, 0x61, 0x64, 0x64, 0x79, 0x03, 0x6e, 0x65, 0x74, 0x00, 0x00, 0x02, 0x00, 0x01, 0x00}
	if bytes.Compare(result, expect) != 0 {
		fmt.Printf("expect: %b\n", expect)
		fmt.Printf("result: %b\n", result)
		fmt.Printf("result: %x\n", result)
		t.Fail()
	}
}

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
