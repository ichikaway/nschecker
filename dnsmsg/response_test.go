package dnsmsg

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckDnsSendIdOk(t *testing.T) {
	var id uint16 = 56474 //valid DNS ID
	message := getDnsResponseBytes()
	result := checkValidDnsId(message, id)
	if result == false {
		t.Fail()
	}
}
func TestCheckDnsSendIdNg(t *testing.T) {
	var id uint16 = 60000 //invalid DNS ID
	message := getDnsResponseBytes()
	result := checkValidDnsId(message, id)
	if result {
		t.Fail()
	}
}

func TestGetNsListFromDnsResponseNoNsServers(t *testing.T) {
	message := []byte{}
	_, err := getNsListFromDnsResponse(message)
	if err == nil {
		t.Fail()
	}
}

func TestGetNsListFromDnsResponseJpServers(t *testing.T) {
	expect := []string{"ns2.bitforest.jp.", "ns1.bitforest.jp."}

	// dns response message byte
	message := getDnsResponseBytes()

	servers, err := getNsListFromDnsResponse(message)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(expect, servers) {
		fmt.Println(expect)
		fmt.Println(servers)
		t.Fail()
	}
}

func TestGetNsListFromDnsResponseJpServersFailCheck(t *testing.T) {
	expect := []string{"ns1.notexistdomain.jp.", "ns2.notexistsdomain.jp."}

	// dns response message byte
	message := getDnsResponseBytes()
	servers, err := getNsListFromDnsResponse(message)
	if err != nil {
		t.Fail()
	}

	if reflect.DeepEqual(expect, servers) {
		fmt.Println(expect)
		fmt.Println(servers)
		t.Fail()
	}
}

func TestAa(t *testing.T) {
	message := getDnsResponseBytes()
	getNsListFromDnsResponse2(message)
}

func getDnsResponseBytes() []byte {
	message := []byte{
		//dns response header
		0b11011100, 0b10011010, //id 16bit  uint16=56474
		0b10000000, 0b00000000, //QR=1 response
		0b00000000, 0b00000001, //Question
		0b00000000, 0b00000000, //Answer
		0b00000000, 0b00000010, //Authority
		0b00000000, 0b00000010, //Additional
		//Question
		0b00001001, 0b01100010, 0b01101001, 0b01110100, 0b01100110, 0b01101111, 0b01110010, 0b01100101, 0b01110011, 0b01110100, //9bitforest
		0b00000010, 0b01101010, 0b01110000, 0b00000000, //2jp\0
		0b00000000, 0b00000010, 0b00000000, 0b00000001, //type=NS, class=inet
		// Authority1
		0b11000000, 0b00001100, //NAME offset pointer 0b00001100(12 byte offset = first byte of Question field)
		0b00000000, 0b00000010, //type=NS
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000001, 0b01010001, 0b10000000, //TTL
		0b00000000, 0b00000110, //length
		0b00000011, 0b01101110, 0b01110011, 0b00110010, //RDATA 3ns2
		0b11000000, 0b00001100, //RDATA offset pointer 0b00001100(12 byte offset = first byte of Question field)
		// Authority2
		0b11000000, 0b00001100, //NAME offset pointer 0b00001100(12 byte offset = first byte of Question field)
		0b00000000, 0b00000010, //type=NS
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000001, 0b01010001, 0b10000000, //TTL
		0b00000000, 0b00000110, //length
		0b00000011, 0b01101110, 0b01110011, 0b00110001, //RDATA 3ns1
		0b11000000, 0b00001100, //RDATA offset pointer 0b00001100(12 byte offset = first byte of Question field)
		//Additional1
		0b11000000, 0b00111100, //NAME
		0b00000000, 0b00000001, //type=A
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000001, 0b01010001, 0b10000000, //TTL
		0b00000000, 0b00000100, //length
		0b00110110, 0b11111000, 0b11011101, 0b00001101, //RDATA
		//Additional2
		0b11000000, 0b00101010, //NAME
		0b00000000, 0b00000001, //type=A
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000001, 0b01010001, 0b10000000, //TTL
		0b00000000, 0b00000100, //length
		0b00110110, 0b10110010, 0b11000101, 0b11100100, //RDATA
	}
	return message
}
