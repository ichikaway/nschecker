package misc

import (
	"fmt"
)

const str string = "\x00\x00\x01\x00\x00\x01\x00\x00\x00\x00\x00\x00\x03www\x04jp0rs\x02co\x02jp\x00\x00\x01\x00\x01"
const str2 string = "\x03www\x04jp0rs\x02co\x02jp\x00\x00\x01\x00\x01"

func checkNullByte() {
	//fmt.Printf("%x", str)
	for i, s := range []byte(str2) {
		fmt.Printf("%d:%x ", i, s)
		if s == 0 {
			fmt.Printf("\nMatch nullbyte")
		}
	}
}
