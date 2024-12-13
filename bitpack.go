package bitpack

import (
	"fmt"
	"math"
	"reflect"
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

		switch switchTypes[T](field.Type) {
		case typeBits:
			size = reflect.New(field.Type).Interface().(bits).Size()
		case typeBool:
			size = 1
		case typeUints, typeInts:
			size = field.Type.Bits()
		}

		typemap[idx] = typeinfo{typ: field.Type, size: size, offset: offset}
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
	lastType := bp.typemap[len(bp.typemap)-1]

	return lastType.offset + lastType.size
}

func (bp *BitPack[T]) Pack(unpacked T) uint64 {
	var pack uint64 = 0

	value := reflect.Indirect(reflect.ValueOf(unpacked))

	for idx, typeinfo := range bp.typemap {
		field := value.Field(idx)
		mask := uint64(math.Pow(2, float64(typeinfo.size))) - 1

		switch switchTypes[T](typeinfo.typ) {
		case typeBits:
			pack |= field.Interface().(bits).Serialize() << uint64(typeinfo.offset)

		case typeBool:
			if field.Bool() {
				pack |= 0b1 << typeinfo.offset
			}

		case typeUints:
			pack |= field.Uint() & mask << uint64(typeinfo.offset)

		case typeInts:
			pack |= uint64(field.Int()) & mask << uint64(typeinfo.offset)
		}
	}

	return pack
}

func (bp *BitPack[T]) Unpack(packed uint64) *T {
	object := new(T)
	value := reflect.Indirect(reflect.ValueOf(object))

	for idx, typeinfo := range bp.typemap {
		mask := uint64(math.Pow(2, float64(typeinfo.size))) - 1
		data := packed >> typeinfo.offset

		if typeinfo.size < 64 {
			data &= mask
		}

		var cast reflect.Value

		switch switchTypes[T](typeinfo.typ) {
		case typeBits:
			cast = reflect.New(typeinfo.typ).Elem()
			cast.Field(0).SetInt(int64(data))

		case typeBool:
			cast = reflect.ValueOf(castBool(data))

		case typeUints, typeInts:
			cast = reflect.ValueOf(castInt(typeinfo.typ, data))
		}

		value.Field(idx).Set(cast)
	}

	return object
}
