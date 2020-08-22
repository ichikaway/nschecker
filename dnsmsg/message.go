package dnsmsg

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"
)

type Type uint16

// CLASS inet
const CLASS_INET uint16 = 1

const (
	TypeNS Type = 2
	//TypeMX Type = 15
)

type Header struct {
	ID     uint16
	QR     byte
	OpCode uint16
	AA     byte
	TC     byte
	RD     byte
	RA     byte
	Z      byte
	RCode  uint16
}

type Header2 struct {
	QD uint16
	AN uint16
	NS uint16
	AR uint16
}

type Question struct {
	Name  string
	Type  Type
	Class uint16
}

// create header.id value
func NewId() uint16 {
	return uint16(rand.Int()) ^ uint16(time.Now().UnixNano())
}

func NewHeader(id uint16) Header {
	return Header{ID: id, RD: 0} // RD=0 to use TLD DNS servers.
}

func NewHeaderPayload(h Header) []byte {
	var buf = new(bytes.Buffer)
	var h_16 byte = 0 //16bit-23bit
	var h_24 byte = 0 //24bit-31bit

	binary.Write(buf, binary.BigEndian, h.ID) //first 16bit

	h_16 = h.QR << 7
	h_16 |= byte(h.OpCode) << (7 - 4)
	h_16 |= h.AA << (7 - 5)
	h_16 |= h.TC << (7 - 6)
	h_16 |= h.RD << (7 - 7)
	binary.Write(buf, binary.BigEndian, h_16)

	h_24 = h.RA << 7
	h_24 |= h.Z << (7 - 3)           //3bit
	h_24 |= byte(h.RCode) << (7 - 7) //4bit
	binary.Write(buf, binary.BigEndian, h_24)

	h2 := Header2{QD: 1}
	binary.Write(buf, binary.BigEndian, h2.QD)
	binary.Write(buf, binary.BigEndian, h2.AN)
	binary.Write(buf, binary.BigEndian, h2.NS)
	binary.Write(buf, binary.BigEndian, h2.AR)
	return buf.Bytes()
}

func NewQuestion(name string, qtype Type) Question {
	q := Question{
		Name:  name,
		Type:  qtype,
		Class: CLASS_INET,
	}
	return q
}

func NewQuestionPayload(q Question) []byte {
	var buf = new(bytes.Buffer)

	// www.example.com -> 3www7example3comに変換
	label := createDomainLabel(q.Name)
	binary.Write(buf, binary.BigEndian, label)
	binary.Write(buf, binary.BigEndian, byte(0)) //end domain name

	binary.Write(buf, binary.BigEndian, q.Type)
	binary.Write(buf, binary.BigEndian, q.Class)
	binary.Write(buf, binary.BigEndian, byte(0)) //end
	return buf.Bytes()
}

// www.example.com -> 3www7example3comに変換
func createDomainLabel(domainame string) []byte {
	var ret []byte
	slice := strings.Split(domainame, ".")
	for _, str := range slice {
		//fmt.Printf("%b", len(str))
		//fmt.Printf("%s\n", str)
		ret = append(ret, byte(len(str)))
		ret = append(ret, []byte(str)...)
	}
	return ret
}
