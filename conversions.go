package bitpack

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

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

	panic("unreachable, BitPack cannot have member")
}
