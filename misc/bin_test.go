package misc

import (
	"fmt"
	"testing"
)

func TestBinCheck(t *testing.T) {
	checkNullByte()
}

func TestShiftBit(t *testing.T) {
	var byteData byte = 0b11000011
	var byteData2 byte = 0b00000010
	var upper byte = byteData & 0b00111111 //上位2bitを落とす
	var under byte = byteData2
	var compressedCounter int16 = int16(upper)<<8 | int16(under) //int16(upper<<8)だとint8を8bitシフトしてから型変換なので8bitしかなくて0になってしまう
	//fmt.Printf("%#016b\n", int16(upper<<8)) //これはバグのケース
	fmt.Printf("%#016b\n", int16(upper)<<8)
	fmt.Printf("%#016b, %#08b, %#08b\n", compressedCounter, upper, under)
}
