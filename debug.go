package bitpack

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"

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
	pack := fmt.Sprintf("%064s", strconv.FormatUint(bp.Pack(value), 2))

	paddings := make([]int, len(bp.typemap))

	for idx := range bp.typemap {
		displayIndex := len(bp.typemap) - idx - 1
		typeinfo := bp.typemap[len(bp.typemap)-idx-1]

		paddings[displayIndex] = PayloadSize - int(typeinfo.offset) - typeinfo.size + idx
		start := PayloadSize - int(typeinfo.offset) - typeinfo.size
		end := start + typeinfo.size

		colors[idx%len(colors)].Print(pack[start:end])

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

		switch switchTypes[T](field.Type) {
		case typeBits:
			v = value.Interface().(bits).Serialize()
		case typeBool:
			v = value.Bool()
		case typeUints:
			v = value.Uint()
		case typeInts:
			v = value.Int()
		case typeFloat32:
			v = value.Float()
		default:
			v = "<unimplemented>"
		}

		printValue(displayIndex, v)
	}
}

func printValue(idx int, value any) {
	colors[idx%len(colors)].Add(color.Concealed).Printf("└ %v\n", value)
}
