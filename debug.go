package bitpack

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var (
	colors = []*color.Color{
		color.New(color.FgCyan),
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgMagenta),
		color.New(color.FgYellow),
	}
)

func (bp *BitPack[T]) Debug(value T) {
	bits := fmt.Sprintf("%064s", strconv.FormatUint(bp.Serialize(&value), 2))

	paddings := make([]int, len(bp.typemap))

	for idx := range bp.typemap {
		displayIndex := len(bp.typemap) - idx - 1
		typeinfo := bp.typemap[len(bp.typemap)-idx-1]

		paddings[displayIndex] = PayloadSize - int(typeinfo.offset) - typeinfo.size + idx
		start := PayloadSize - int(typeinfo.offset) - typeinfo.size
		end := start + typeinfo.size

		colors[idx%len(colors)].Print(bits[start:end])

		fmt.Print(" ")
	}

	fmt.Println()

	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)

	for idx := range bp.typemap {
		displayIndex := len(bp.typemap) - idx - 1
		field := typ.Field(idx)
		value := val.Field(idx)

		colorIdx := 0

		for col := range paddings[idx] {
			if slices.Contains(paddings, col) {
				colors[colorIdx%len(colors)].Add(color.Concealed).Print("│")
				colorIdx += 1
			} else {
				fmt.Print(" ")
			}
		}

		var v any

		switch {
		case strings.HasPrefix(field.Type.String(), "bitpack.Bit"):
			v = value.Interface().(Arbitrary).Serialize()

		default:
			switch field.Type.Kind() {
			case reflect.Bool:
				v = value.Bool()

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				v = value.Uint()

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v = value.Int()

			default:
				panic(fmt.Sprintf("%T contains a %s field, which is not supported by BitPacks", *new(T), field.Type.Kind()))
			}
		}

		printValue(displayIndex, v)
	}
}

func printValue(idx int, value any) {
	colors[idx%len(colors)].Add(color.Concealed).Printf("└ %v\n", value)
}
