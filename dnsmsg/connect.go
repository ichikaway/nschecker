package dnsmsg

import (
	"net"
	"time"
)

// send UDP Packet and get response.
func Send(address string, message []byte) ([]byte, error) {
	timeout := 10 * time.Second
	d := net.Dialer{Timeout: timeout}
	conn, err := d.Dial("udp4", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	//fmt.Println("Sending to server")
	_, err = conn.Write(message)
	if err != nil {
		return nil, err
	}

	//fmt.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Received(x): %x\n", buffer[:length])
	//fmt.Printf("Received(s): %s\n", string(buffer[:length]))
	return buffer[:length], nil
}
