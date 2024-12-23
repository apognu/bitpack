package main

import (
	"fmt"
	"strconv"

	"github.com/apognu/bitpack"
)

type MyType struct {
	A bitpack.Bit1
	B uint8
	C bool
	D int16
	E bitpack.Bit6
	F float32
}

var (
	bp = bitpack.NewBitPack[MyType]()
)

func main() {
	t := MyType{
		A: bitpack.NewBit1(0b1),
		B: 42,
		C: true,
		D: -9000,
		E: bitpack.NewBit6(0b111001),
		F: 42.9001,
	}

	pack := bp.Pack(t)

	fmt.Println("PACK  ", strconv.FormatUint(pack, 2))

	unpack := bp.Unpack(pack)

	fmt.Printf("STRUCT %#v\n", unpack)

	bp.Debug(t)
}
