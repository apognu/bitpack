package bitpack

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"
)

type types int

const (
	typeBits types = iota
	typeBool
	typeUints
	typeInts
)

const (
	BitsPrefix = "bitpack.Bit"
)

func switchTypes[BP any](typ reflect.Type) types {
	switch {
	case strings.HasPrefix(typ.String(), BitsPrefix):
		return typeBits

	default:
		switch typ.Kind() {
		case reflect.Bool:
			return typeBool

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return typeUints

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return typeInts

		default:
			panic(fmt.Sprintf("%T contains a %s field, which is not supported by BitPacks", *new(BP), typ.Kind()))
		}
	}
}

func castBool(value uint64) bool {
	return value == 1
}

func castInt[T constraints.Integer](typ reflect.Type, value T) any {
	switch typ.Kind() {
	case reflect.Uint:
		return uint(value)
	case reflect.Int:
		return int(value)
	case reflect.Uint8:
		return uint8(value)
	case reflect.Int8:
		return int8(value)
	case reflect.Uint16:
		return uint16(value)
	case reflect.Int16:
		return int16(value)
	case reflect.Uint32:
		return uint32(value)
	case reflect.Int32:
		return int32(value)
	case reflect.Uint64:
		return uint64(value)
	case reflect.Int64:
		return int64(value)
	}

	panic(fmt.Sprintf("unreachable, BitPack cannot have %T member", value))
}
