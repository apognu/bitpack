package bitpack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrongBackingType(t *testing.T) {
	assert.Panics(t, func() { NewBitPack[int]() }, "wrong backing type was accepted")
}

func TestOversizedStruct(t *testing.T) {
	type twoInt64 struct {
		A int64
		B int64
	}

	assert.Panics(t, func() { NewBitPack[twoInt64]() }, "oversized struct was accepted")
}

func TestWrongMemberType(t *testing.T) {
	type wrongType struct {
		A int32
		B string
	}

	assert.Panics(t, func() { NewBitPack[wrongType]() }, "struct with invalid member was accepted")
}

func TestFourInt16(t *testing.T) {
	type fourInt16 struct {
		A int16
		B int16
		C int16
		D int16
	}

	bp := NewBitPack[fourInt16]()
	v := fourInt16{1, 2, 3, 4}

	assert.Equal(t, uint64(0b100000000000000001100000000000000100000000000000001), bp.Pack(v))
	assert.Equal(t, v, *bp.Unpack(bp.Pack(v)))
}

func TestOneInt64(t *testing.T) {
	type oneInt64 struct {
		A int64
	}

	bp := NewBitPack[oneInt64]()
	v := oneInt64{42}

	assert.Equal(t, uint64(0b101010), bp.Pack(v))
	assert.Equal(t, v, *bp.Unpack(bp.Pack(v)))
}

func TestAllUnsignedTypes(t *testing.T) {
	type allUnsignedTypes struct {
		A uint32
		B uint16
		C uint8
		D bool
	}

	bp := NewBitPack[allUnsignedTypes]()
	v := allUnsignedTypes{200000, 9000, 42, true}

	assert.Equal(t, uint64(0b100101010001000110010100000000000000000110000110101000000), bp.Pack(v))
	assert.Equal(t, v, *bp.Unpack(bp.Pack(v)))
}

func TestAllSignedTypes(t *testing.T) {
	type allSignedTypes struct {
		A int32
		B int16
		C int8
		D bool
	}

	bp := NewBitPack[allSignedTypes]()
	v := allSignedTypes{-200000, 9000, -42, true}

	assert.Equal(t, uint64(0b111010110001000110010100011111111111111001111001011000000), bp.Pack(v))
	assert.Equal(t, v, *bp.Unpack(bp.Pack(v)))
}

func TestMixed(t *testing.T) {
	type mixed struct {
		A int8
		B bool
		C uint32
		D int16
	}

	bp := NewBitPack[mixed]()
	v := mixed{-128, true, 9000, -1}

	assert.Equal(t, uint64(0b111111111111111100000000000000000010001100101000110000000), bp.Pack(v))
	assert.Equal(t, v, *bp.Unpack(bp.Pack(v)))
}

func TestBooleans(t *testing.T) {
	type tenBooleans struct {
		A bool
		B bool
		C bool
		D bool
		E bool
		F bool
		G bool
		H bool
		I bool
		J bool
	}

	bp := NewBitPack[tenBooleans]()
	v := tenBooleans{true, false, false, true, true, false, false, true, false, true}

	assert.Equal(t, uint64(0b1010011001), bp.Pack(v))
	assert.Equal(t, v, *bp.Unpack(bp.Pack(v)))
}

func TestArbitraryBits(t *testing.T) {
	type tenVarying struct {
		One    Bit1
		Two    Bit2
		Four   Bit4
		Fourty Bit40
	}

	bp := NewBitPack[tenVarying]()
	v := tenVarying{NewBit1(0b1), NewBit2(0b10), NewBit4(0b1010), NewBit40(0b111111)}

	assert.Equal(t, uint64(0b1111111010101), bp.Pack(v))
	assert.Equal(t, v, *bp.Unpack(bp.Pack(v)))
}

func TestArbitraryBoundsCheck(t *testing.T) {
	assert.NotPanics(t, func() { NewBit4(0b1111) })
	assert.Panics(t, func() { NewBit4(0b11111) })
	assert.NotPanics(t, func() { NewBit12(0b100000000000) })
	assert.Panics(t, func() { NewBit12(0b1000000000000) })
}
