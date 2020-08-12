package dnsmsg

import (
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
func TestMessage(t *testing.T) {
	id := uint16(60000)
	h := NewHeader(id)
	req := NewHeaderPayload(h)
	//fmt.Printf("%v\n", h)

	// 構造体をバイナリにする
	//b := bytes.NewBuffer([]byte(""))
	//binary.Write(b, binary.BigEndian, h)
	//binary.Write(b, binary.LittleEndian, &a)
	/*
	fmt.Printf("hex: %x\n",b.Bytes())
	fmt.Printf("hex: %x\n",b)
	fmt.Printf("bin: %b\n",b.Bytes())
	fmt.Printf("len: %d\n",len(b.Bytes()))
	*/
	fmt.Printf("bin: %b\n", req)
	// ea60 0100 0001 0000 0000 0000

	//hoge := []byte{0xea, 0x60}
	//fmt.Printf("hoge: %d\n", binary.BigEndian.Uint16(hoge))

	q := NewQuestion("vaddy.net", TypeNS)
	req2 := NewQuestionPayload(q)
	fmt.Printf("bin question: %b\n", req2)
}


