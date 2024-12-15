package bitpack

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchTypes(t *testing.T) {
	tt := []struct {
		Type  types
		Value any
	}{
		{typeBits, NewBit1(0)},
		{typeBits, NewBit16(0)},
		{typeBool, true},
		{typeBool, false},
		{typeUints, uint(42)},
		{typeInts, int(42)},
		{typeUints, uint8(42)},
		{typeInts, int8(42)},
		{typeUints, uint16(42)},
		{typeInts, int16(42)},
		{typeUints, uint32(42)},
		{typeInts, int32(42)},
		{typeUints, uint64(42)},
		{typeInts, int64(42)},
	}

	for _, spec := range tt {
		assert.Equal(t, spec.Type, switchTypes[any](reflect.TypeOf(spec.Value)))
	}
}

func TestUnsuppotedSwitchTypes(t *testing.T) {
	assert.Panics(t, func() { switchTypes[any](reflect.TypeOf("string")) })
}

func TestCastBool(t *testing.T) {
	assert.Equal(t, true, castBool(1))
	assert.Equal(t, false, castBool(0))
}

func TestCastInt(t *testing.T) {
	tt := []struct {
		Type  reflect.Type
		Value int
	}{
		{reflect.TypeOf(uint(0)), 42},
		{reflect.TypeOf(int(0)), 42},
		{reflect.TypeOf(uint8(0)), 42},
		{reflect.TypeOf(int8(0)), 42},
		{reflect.TypeOf(uint16(0)), 42},
		{reflect.TypeOf(int16(0)), 42},
		{reflect.TypeOf(uint32(0)), 42},
		{reflect.TypeOf(int32(0)), 42},
		{reflect.TypeOf(uint64(0)), 42},
		{reflect.TypeOf(int64(0)), 42},
	}

	for _, spec := range tt {
		out := castInt(spec.Type, spec.Value)

		assert.Equal(t, spec.Type, reflect.TypeOf(out))
		assert.EqualValues(t, spec.Value, out)
	}
}

func TestUnreachableCastInt(t *testing.T) {
	assert.Panics(t, func() { castInt(reflect.TypeOf("string"), 12) })
}
