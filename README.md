# bitpack

`bitpack` handles packing a struct into a single `uint64` that can easily be saved as a single value or used with atomics.

It reads the structure of the struct it is initialized on to calculate offsets used when packing and unpacking. Since it is ultimately encoded into a 64-bits integer, the sum of all members cannot exceed 64 bits. As of now, it can pack structs composed of the following types:

 - All Go integers types
 - Booleans
 - Arbitrary-length unsigned integers provided by this package's `Bit[1-64]` structs

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/apognu/bitpack"
)

// Create a regular struct with the supported types
type MyType struct {
	A bitpack.Bit1
	B uint8
	C bool
	D int16
	E bitpack.Bit38
}

var (
	// Initialize a bitpack struct on that type
	bp = bitpack.NewBitPack[MyType]()
)

func main() {
	t := MyType{
		A: bitpack.NewBit1(0b1),
		B: 42,
		C: true,
		D: -9000,
		E: bitpack.NewBit38(0b111001),
	}

	// Pack the struct into a uint64
	pack := bp.Pack(t)

	fmt.Println("PACK  ", strconv.FormatUint(pack, 2))

	// Unpack a uint64 into the initial struct
	fmt.Printf("STRUCT %#v\n", bp.Unpack(pack))

	bp.Debug(t)
}
```

```shell
$ go run example/main.go
PACK   11100111011100110110001001010101
STRUCT &main.MyType{A:bitpack.Bit1{Value:1}, B:0x2a, C:true, D:-9000, E:bitpack.Bit38{Value:57}}
00000000000000000000000000000000111001 1101110011011000 1 00101010 1
│                                      │                │ │        └ 1
│                                      │                │ └ 42
│                                      │                └ true
│                                      └ -9000
└ 57
```
