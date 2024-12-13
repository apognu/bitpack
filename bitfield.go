package bitpack

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

//go:generate go run gen/main.go

const (
	PayloadSize = 64
)

type typeinfo struct {
	typ    reflect.Type
	size   int
	offset int
}

type BitPack[T any] struct {
	typemap []typeinfo
}

func NewBitPack[T any]() *BitPack[T] {
	typ := reflect.TypeOf((*T)(nil))

	if typ.Elem().Kind() != reflect.Struct {
		panic(fmt.Sprintf("BitPacks can only built from structs, %T is %s", *new(T), typ.Elem().Kind()))
	}

	typemap := make([]typeinfo, typ.Elem().NumField())
	offset := 0

	for idx := 0; idx < typ.Elem().NumField(); idx++ {
		field := typ.Elem().Field(idx)
		size := 0

		switch {
		case strings.HasPrefix(field.Type.String(), "bitpack.Bit"):
			arbSize := reflect.New(field.Type).Interface().(Arbitrary)

			size = arbSize.Size()
			typemap[idx] = typeinfo{typ: field.Type, size: size, offset: offset}

		default:
			switch field.Type.Kind() {
			case reflect.Bool:
				size = 1
				typemap[idx] = typeinfo{typ: field.Type, size: size, offset: offset}

			case reflect.Int, reflect.Uint, reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
				size = field.Type.Bits()
				typemap[idx] = typeinfo{typ: field.Type, size: size, offset: offset}

			default:
				panic(fmt.Sprintf("%T contains a %s field, which is not supported by BitPacks", *new(T), field.Type.Kind()))
			}
		}

		offset += size
	}

	if offset > PayloadSize {
		panic(fmt.Sprintf("%T is %d bits long, but BitPacks can only store %d bits", *new(T), offset, PayloadSize))
	}

	return &BitPack[T]{
		typemap: typemap,
	}
}

func (bp BitPack[T]) Size() int {
	return int(bp.typemap[len(bp.typemap)-1].offset) + bp.typemap[len(bp.typemap)-1].size
}

func (bp *BitPack[T]) Serialize(ser *T) uint64 {
	var bits uint64 = 0

	value := reflect.Indirect(reflect.ValueOf(ser))
	offset := 0

	for idx, typeinfo := range bp.typemap {
		field := value.Field(idx)
		mask := uint64(math.Pow(2, float64(typeinfo.size))) - 1

		switch {
		case strings.HasPrefix(typeinfo.typ.String(), "bitpack.Bit"):
			arbSize := field.Interface().(Arbitrary)

			bits = bits | arbSize.Serialize()<<offset

		default:
			switch typeinfo.typ.Kind() {
			case reflect.Bool:
				if field.Bool() {
					bits = bits | 0b1<<offset
				} else {
					bits = bits | 0b0<<offset
				}

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				bits = bits | field.Uint()&mask<<offset

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				bits = bits | uint64(field.Int())&mask<<offset

			default:
				panic("unreachable, BitPack cannot exist with the provided member")
			}
		}

		offset += typeinfo.size
	}

	return bits
}

func (bp *BitPack[T]) Deserialize(bits uint64) *T {
	output := new(T)
	value := reflect.Indirect(reflect.ValueOf(output))

	for idx, typeinfo := range bp.typemap {
		mask := uint64(math.Pow(2, float64(typeinfo.size))) - 1
		data := bits >> typeinfo.offset

		if typeinfo.size < 64 {
			data &= mask
		}

		var cast any

		switch {
		case strings.HasPrefix(typeinfo.typ.String(), "bitpack.Bit"):
			arbSize := reflect.New(typeinfo.typ).Elem()
			arbSize.Field(0).SetInt(int64(data))

			value.Field(idx).Set(arbSize)

		default:
			switch typeinfo.typ.Kind() {
			case reflect.Bool:
				cast = castBool(data)
				value.Field(idx).Set(reflect.ValueOf(cast))

			case reflect.Int, reflect.Uint, reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
				cast = castInt(typeinfo.typ, data)
				value.Field(idx).Set(reflect.ValueOf(cast))

			default:
				panic("unreachable, BitPack cannot exist with provided members")
			}
		}
	}

	return output
}
