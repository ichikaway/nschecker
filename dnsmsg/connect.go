package dnsmsg

import (
	"fmt"
	"net"
)

func Send(address string, message []byte) []byte {
	conn, err := net.Dial("udp4", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Sending to server")
	_, err = conn.Write(message)
	if err != nil {
		panic(err)
	}

	fmt.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Received(x): %x\n", buffer[:length])
	//fmt.Printf("Received(s): %s\n", string(buffer[:length]))
	return buffer[:length]
}
