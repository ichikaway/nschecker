package dnsmsg

import (
	"fmt"
	"reflect"
	"testing"
)

//cover test case for dnsmessage with compression domains and nocompression domains
func TestGetNsListFromDnsResponseNetServers(t *testing.T) {
	expect := []string{
		"ns-700.awsdns-23.net.",
		"ns-457.awsdns-57.com.",
		"ns-1151.awsdns-15.org.",
		"ns-1908.awsdns-46.co.uk.",
	}

	// dns response message byte
	message := getDnsResponseBytesNoCompression()

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

func TestGetNsListFromDnsResponseNetServersFailCheck(t *testing.T) {
	expect := []string{"ns-700.awsdns-23.net.", "ns1.notexistdomain.jp.", "ns2.notexistsdomain.jp."}

	// dns response message byte
	message := getDnsResponseBytesNoCompression()
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

func getDnsResponseBytesNoCompression() []byte {
	message := []byte{
		//dns response header
		0b11011100, 0b10011010, //id 16bit  uint16=56474
		0b10000000, 0b00000000, //QR=1 response
		0b00000000, 0b00000001, //Question
		0b00000000, 0b00000000, //Answer
		0b00000000, 0b00000100, //Authority 4records
		0b00000000, 0b00000001, //Additional

		//Question
		0b00000101, 0b01110110, 0b01100001, 0b01100100, 0b01100100, 0b01111001, //5vaddy
		0b00000011, 0b01101110, 0b01100101, 0b01110100, 0b00000000, //3net\0
		0b00000000, 0b00000010, 0b00000000, 0b00000001, //type=NS, class=inet

		// Authority1
		0b11000000, 0b00001100, //NAME offset pointer 0b00001100(12 byte offset = first byte of Question field)
		0b00000000, 0b00000010, //type=NS
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000010, 0b10100011, 0b00000000, //TTL
		0b00000000, 0b00010011, //length
		//RDATA ns-700.awsdns-23.net
		0b00000110, 0b01101110, 0b01110011, 0b00101101, 0b00110111, 0b00110000, 0b00110000, //6ns-700
		0b00001001, 0b01100001, 0b01110111, 0b01110011, 0b01100100, 0b01101110, 0b01110011, 0b00101101, 0b00110010, 0b00110011, //9awsdns-23
		0b11000000, 0b00010010, //offset pointer 0b00010010(18 byte offset = first byte of Question field)

		// Authority2
		0b11000000, 0b00001100, //NAME offset pointer 0b00001100(12 byte offset = first byte of Question field)
		0b00000000, 0b00000010, //type=NS
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000010, 0b10100011, 0b00000000, //TTL
		0b00000000, 0b00010110, //length
		//RDATA ns-457.awsdns-57.com
		0b00000110, 0b01101110, 0b01110011, 0b00101101, 0b00110100, 0b00110101, 0b00110111, //6ns-457
		0b00001001, 0b01100001, 0b01110111, 0b01110011, 0b01100100, 0b01101110, 0b01110011, 0b00101101, 0b00110101, 0b00110111, //9awsdns-57
		0b00000011, 0b01100011, 0b01101111, 0b01101101, 0b00000000, //3com\0

		// Authority3
		0b11000000, 0b00001100, //NAME offset pointer 0b00001100(12 byte offset = first byte of Question field)
		0b00000000, 0b00000010, //type=NS
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000010, 0b10100011, 0b00000000, //TTL
		0b00000000, 0b00010111, //length
		//RDATA ns-1151.awsdns-15.org
		0b00000111, 0b01101110, 0b01110011, 0b00101101, 0b00110001, 0b00110001, 0b00110101, 0b00110001, //7ns-1151
		0b00001001, 0b01100001, 0b01110111, 0b01110011, 0b01100100, 0b01101110, 0b01110011, 0b00101101, 0b00110001, 0b00110101, //9awsdns-15
		0b00000011, 0b01101111, 0b01110010, 0b01100111, 0b00000000, //3org\0

		// Authority4
		0b11000000, 0b00001100, //NAME offset pointer 0b00001100(12 byte offset = first byte of Question field)
		0b00000000, 0b00000010, //type=NS
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000010, 0b10100011, 0b00000000, //TTL
		0b00000000, 0b00011001, //length
		//RDATA ns-1908.awsdns-46.co.uk
		0b00000111, 0b01101110, 0b01110011, 0b00101101, 0b00110001, 0b00111001, 0b00110000, 0b00111000, //7ns-1908
		0b00001001, 0b01100001, 0b01110111, 0b01110011, 0b01100100, 0b01101110, 0b01110011, 0b00101101, 0b00110100, 0b00110110, //9awsdns-46
		0b00000010, 0b01100011, 0b01101111, //2co
		0b00000010, 0b01110101, 0b01101011, 0b00000000, //2uk\0

		//Additional
		0b11000000, 0b00100111, //NAME offset pointer
		0b00000000, 0b00000001, //type=A
		0b00000000, 0b00000001, //class=inet
		0b00000000, 0b00000010, 0b10100011, 0b00000000, //TTL
		0b00000000, 0b00000100, //length
		0b11001101, 0b11111011, 0b11000010, 0b10111100, //RDATA
	}
	return message
}
