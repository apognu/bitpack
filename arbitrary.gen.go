package bitpack

import (
	"fmt"
	"math"
)


type Bit1 struct {
  Value int
}

func NewBit1(value int) Bit1 {
  if value >= int(math.Pow(2, float64(1))) {
    panic(fmt.Sprintf("number %d cannot be stored in 1 bits", value))
  }

  return Bit1{Value: value}
}

func (Bit1) Size() int {
  return 1
}

func (b Bit1) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit2 struct {
  Value int
}

func NewBit2(value int) Bit2 {
  if value >= int(math.Pow(2, float64(2))) {
    panic(fmt.Sprintf("number %d cannot be stored in 2 bits", value))
  }

  return Bit2{Value: value}
}

func (Bit2) Size() int {
  return 2
}

func (b Bit2) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit3 struct {
  Value int
}

func NewBit3(value int) Bit3 {
  if value >= int(math.Pow(2, float64(3))) {
    panic(fmt.Sprintf("number %d cannot be stored in 3 bits", value))
  }

  return Bit3{Value: value}
}

func (Bit3) Size() int {
  return 3
}

func (b Bit3) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit4 struct {
  Value int
}

func NewBit4(value int) Bit4 {
  if value >= int(math.Pow(2, float64(4))) {
    panic(fmt.Sprintf("number %d cannot be stored in 4 bits", value))
  }

  return Bit4{Value: value}
}

func (Bit4) Size() int {
  return 4
}

func (b Bit4) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit5 struct {
  Value int
}

func NewBit5(value int) Bit5 {
  if value >= int(math.Pow(2, float64(5))) {
    panic(fmt.Sprintf("number %d cannot be stored in 5 bits", value))
  }

  return Bit5{Value: value}
}

func (Bit5) Size() int {
  return 5
}

func (b Bit5) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit6 struct {
  Value int
}

func NewBit6(value int) Bit6 {
  if value >= int(math.Pow(2, float64(6))) {
    panic(fmt.Sprintf("number %d cannot be stored in 6 bits", value))
  }

  return Bit6{Value: value}
}

func (Bit6) Size() int {
  return 6
}

func (b Bit6) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit7 struct {
  Value int
}

func NewBit7(value int) Bit7 {
  if value >= int(math.Pow(2, float64(7))) {
    panic(fmt.Sprintf("number %d cannot be stored in 7 bits", value))
  }

  return Bit7{Value: value}
}

func (Bit7) Size() int {
  return 7
}

func (b Bit7) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit8 struct {
  Value int
}

func NewBit8(value int) Bit8 {
  if value >= int(math.Pow(2, float64(8))) {
    panic(fmt.Sprintf("number %d cannot be stored in 8 bits", value))
  }

  return Bit8{Value: value}
}

func (Bit8) Size() int {
  return 8
}

func (b Bit8) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit9 struct {
  Value int
}

func NewBit9(value int) Bit9 {
  if value >= int(math.Pow(2, float64(9))) {
    panic(fmt.Sprintf("number %d cannot be stored in 9 bits", value))
  }

  return Bit9{Value: value}
}

func (Bit9) Size() int {
  return 9
}

func (b Bit9) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit10 struct {
  Value int
}

func NewBit10(value int) Bit10 {
  if value >= int(math.Pow(2, float64(10))) {
    panic(fmt.Sprintf("number %d cannot be stored in 10 bits", value))
  }

  return Bit10{Value: value}
}

func (Bit10) Size() int {
  return 10
}

func (b Bit10) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit11 struct {
  Value int
}

func NewBit11(value int) Bit11 {
  if value >= int(math.Pow(2, float64(11))) {
    panic(fmt.Sprintf("number %d cannot be stored in 11 bits", value))
  }

  return Bit11{Value: value}
}

func (Bit11) Size() int {
  return 11
}

func (b Bit11) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit12 struct {
  Value int
}

func NewBit12(value int) Bit12 {
  if value >= int(math.Pow(2, float64(12))) {
    panic(fmt.Sprintf("number %d cannot be stored in 12 bits", value))
  }

  return Bit12{Value: value}
}

func (Bit12) Size() int {
  return 12
}

func (b Bit12) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit13 struct {
  Value int
}

func NewBit13(value int) Bit13 {
  if value >= int(math.Pow(2, float64(13))) {
    panic(fmt.Sprintf("number %d cannot be stored in 13 bits", value))
  }

  return Bit13{Value: value}
}

func (Bit13) Size() int {
  return 13
}

func (b Bit13) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit14 struct {
  Value int
}

func NewBit14(value int) Bit14 {
  if value >= int(math.Pow(2, float64(14))) {
    panic(fmt.Sprintf("number %d cannot be stored in 14 bits", value))
  }

  return Bit14{Value: value}
}

func (Bit14) Size() int {
  return 14
}

func (b Bit14) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit15 struct {
  Value int
}

func NewBit15(value int) Bit15 {
  if value >= int(math.Pow(2, float64(15))) {
    panic(fmt.Sprintf("number %d cannot be stored in 15 bits", value))
  }

  return Bit15{Value: value}
}

func (Bit15) Size() int {
  return 15
}

func (b Bit15) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit16 struct {
  Value int
}

func NewBit16(value int) Bit16 {
  if value >= int(math.Pow(2, float64(16))) {
    panic(fmt.Sprintf("number %d cannot be stored in 16 bits", value))
  }

  return Bit16{Value: value}
}

func (Bit16) Size() int {
  return 16
}

func (b Bit16) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit17 struct {
  Value int
}

func NewBit17(value int) Bit17 {
  if value >= int(math.Pow(2, float64(17))) {
    panic(fmt.Sprintf("number %d cannot be stored in 17 bits", value))
  }

  return Bit17{Value: value}
}

func (Bit17) Size() int {
  return 17
}

func (b Bit17) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit18 struct {
  Value int
}

func NewBit18(value int) Bit18 {
  if value >= int(math.Pow(2, float64(18))) {
    panic(fmt.Sprintf("number %d cannot be stored in 18 bits", value))
  }

  return Bit18{Value: value}
}

func (Bit18) Size() int {
  return 18
}

func (b Bit18) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit19 struct {
  Value int
}

func NewBit19(value int) Bit19 {
  if value >= int(math.Pow(2, float64(19))) {
    panic(fmt.Sprintf("number %d cannot be stored in 19 bits", value))
  }

  return Bit19{Value: value}
}

func (Bit19) Size() int {
  return 19
}

func (b Bit19) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit20 struct {
  Value int
}

func NewBit20(value int) Bit20 {
  if value >= int(math.Pow(2, float64(20))) {
    panic(fmt.Sprintf("number %d cannot be stored in 20 bits", value))
  }

  return Bit20{Value: value}
}

func (Bit20) Size() int {
  return 20
}

func (b Bit20) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit21 struct {
  Value int
}

func NewBit21(value int) Bit21 {
  if value >= int(math.Pow(2, float64(21))) {
    panic(fmt.Sprintf("number %d cannot be stored in 21 bits", value))
  }

  return Bit21{Value: value}
}

func (Bit21) Size() int {
  return 21
}

func (b Bit21) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit22 struct {
  Value int
}

func NewBit22(value int) Bit22 {
  if value >= int(math.Pow(2, float64(22))) {
    panic(fmt.Sprintf("number %d cannot be stored in 22 bits", value))
  }

  return Bit22{Value: value}
}

func (Bit22) Size() int {
  return 22
}

func (b Bit22) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit23 struct {
  Value int
}

func NewBit23(value int) Bit23 {
  if value >= int(math.Pow(2, float64(23))) {
    panic(fmt.Sprintf("number %d cannot be stored in 23 bits", value))
  }

  return Bit23{Value: value}
}

func (Bit23) Size() int {
  return 23
}

func (b Bit23) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit24 struct {
  Value int
}

func NewBit24(value int) Bit24 {
  if value >= int(math.Pow(2, float64(24))) {
    panic(fmt.Sprintf("number %d cannot be stored in 24 bits", value))
  }

  return Bit24{Value: value}
}

func (Bit24) Size() int {
  return 24
}

func (b Bit24) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit25 struct {
  Value int
}

func NewBit25(value int) Bit25 {
  if value >= int(math.Pow(2, float64(25))) {
    panic(fmt.Sprintf("number %d cannot be stored in 25 bits", value))
  }

  return Bit25{Value: value}
}

func (Bit25) Size() int {
  return 25
}

func (b Bit25) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit26 struct {
  Value int
}

func NewBit26(value int) Bit26 {
  if value >= int(math.Pow(2, float64(26))) {
    panic(fmt.Sprintf("number %d cannot be stored in 26 bits", value))
  }

  return Bit26{Value: value}
}

func (Bit26) Size() int {
  return 26
}

func (b Bit26) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit27 struct {
  Value int
}

func NewBit27(value int) Bit27 {
  if value >= int(math.Pow(2, float64(27))) {
    panic(fmt.Sprintf("number %d cannot be stored in 27 bits", value))
  }

  return Bit27{Value: value}
}

func (Bit27) Size() int {
  return 27
}

func (b Bit27) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit28 struct {
  Value int
}

func NewBit28(value int) Bit28 {
  if value >= int(math.Pow(2, float64(28))) {
    panic(fmt.Sprintf("number %d cannot be stored in 28 bits", value))
  }

  return Bit28{Value: value}
}

func (Bit28) Size() int {
  return 28
}

func (b Bit28) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit29 struct {
  Value int
}

func NewBit29(value int) Bit29 {
  if value >= int(math.Pow(2, float64(29))) {
    panic(fmt.Sprintf("number %d cannot be stored in 29 bits", value))
  }

  return Bit29{Value: value}
}

func (Bit29) Size() int {
  return 29
}

func (b Bit29) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit30 struct {
  Value int
}

func NewBit30(value int) Bit30 {
  if value >= int(math.Pow(2, float64(30))) {
    panic(fmt.Sprintf("number %d cannot be stored in 30 bits", value))
  }

  return Bit30{Value: value}
}

func (Bit30) Size() int {
  return 30
}

func (b Bit30) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit31 struct {
  Value int
}

func NewBit31(value int) Bit31 {
  if value >= int(math.Pow(2, float64(31))) {
    panic(fmt.Sprintf("number %d cannot be stored in 31 bits", value))
  }

  return Bit31{Value: value}
}

func (Bit31) Size() int {
  return 31
}

func (b Bit31) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit32 struct {
  Value int
}

func NewBit32(value int) Bit32 {
  if value >= int(math.Pow(2, float64(32))) {
    panic(fmt.Sprintf("number %d cannot be stored in 32 bits", value))
  }

  return Bit32{Value: value}
}

func (Bit32) Size() int {
  return 32
}

func (b Bit32) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit33 struct {
  Value int
}

func NewBit33(value int) Bit33 {
  if value >= int(math.Pow(2, float64(33))) {
    panic(fmt.Sprintf("number %d cannot be stored in 33 bits", value))
  }

  return Bit33{Value: value}
}

func (Bit33) Size() int {
  return 33
}

func (b Bit33) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit34 struct {
  Value int
}

func NewBit34(value int) Bit34 {
  if value >= int(math.Pow(2, float64(34))) {
    panic(fmt.Sprintf("number %d cannot be stored in 34 bits", value))
  }

  return Bit34{Value: value}
}

func (Bit34) Size() int {
  return 34
}

func (b Bit34) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit35 struct {
  Value int
}

func NewBit35(value int) Bit35 {
  if value >= int(math.Pow(2, float64(35))) {
    panic(fmt.Sprintf("number %d cannot be stored in 35 bits", value))
  }

  return Bit35{Value: value}
}

func (Bit35) Size() int {
  return 35
}

func (b Bit35) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit36 struct {
  Value int
}

func NewBit36(value int) Bit36 {
  if value >= int(math.Pow(2, float64(36))) {
    panic(fmt.Sprintf("number %d cannot be stored in 36 bits", value))
  }

  return Bit36{Value: value}
}

func (Bit36) Size() int {
  return 36
}

func (b Bit36) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit37 struct {
  Value int
}

func NewBit37(value int) Bit37 {
  if value >= int(math.Pow(2, float64(37))) {
    panic(fmt.Sprintf("number %d cannot be stored in 37 bits", value))
  }

  return Bit37{Value: value}
}

func (Bit37) Size() int {
  return 37
}

func (b Bit37) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit38 struct {
  Value int
}

func NewBit38(value int) Bit38 {
  if value >= int(math.Pow(2, float64(38))) {
    panic(fmt.Sprintf("number %d cannot be stored in 38 bits", value))
  }

  return Bit38{Value: value}
}

func (Bit38) Size() int {
  return 38
}

func (b Bit38) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit39 struct {
  Value int
}

func NewBit39(value int) Bit39 {
  if value >= int(math.Pow(2, float64(39))) {
    panic(fmt.Sprintf("number %d cannot be stored in 39 bits", value))
  }

  return Bit39{Value: value}
}

func (Bit39) Size() int {
  return 39
}

func (b Bit39) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit40 struct {
  Value int
}

func NewBit40(value int) Bit40 {
  if value >= int(math.Pow(2, float64(40))) {
    panic(fmt.Sprintf("number %d cannot be stored in 40 bits", value))
  }

  return Bit40{Value: value}
}

func (Bit40) Size() int {
  return 40
}

func (b Bit40) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit41 struct {
  Value int
}

func NewBit41(value int) Bit41 {
  if value >= int(math.Pow(2, float64(41))) {
    panic(fmt.Sprintf("number %d cannot be stored in 41 bits", value))
  }

  return Bit41{Value: value}
}

func (Bit41) Size() int {
  return 41
}

func (b Bit41) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit42 struct {
  Value int
}

func NewBit42(value int) Bit42 {
  if value >= int(math.Pow(2, float64(42))) {
    panic(fmt.Sprintf("number %d cannot be stored in 42 bits", value))
  }

  return Bit42{Value: value}
}

func (Bit42) Size() int {
  return 42
}

func (b Bit42) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit43 struct {
  Value int
}

func NewBit43(value int) Bit43 {
  if value >= int(math.Pow(2, float64(43))) {
    panic(fmt.Sprintf("number %d cannot be stored in 43 bits", value))
  }

  return Bit43{Value: value}
}

func (Bit43) Size() int {
  return 43
}

func (b Bit43) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit44 struct {
  Value int
}

func NewBit44(value int) Bit44 {
  if value >= int(math.Pow(2, float64(44))) {
    panic(fmt.Sprintf("number %d cannot be stored in 44 bits", value))
  }

  return Bit44{Value: value}
}

func (Bit44) Size() int {
  return 44
}

func (b Bit44) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit45 struct {
  Value int
}

func NewBit45(value int) Bit45 {
  if value >= int(math.Pow(2, float64(45))) {
    panic(fmt.Sprintf("number %d cannot be stored in 45 bits", value))
  }

  return Bit45{Value: value}
}

func (Bit45) Size() int {
  return 45
}

func (b Bit45) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit46 struct {
  Value int
}

func NewBit46(value int) Bit46 {
  if value >= int(math.Pow(2, float64(46))) {
    panic(fmt.Sprintf("number %d cannot be stored in 46 bits", value))
  }

  return Bit46{Value: value}
}

func (Bit46) Size() int {
  return 46
}

func (b Bit46) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit47 struct {
  Value int
}

func NewBit47(value int) Bit47 {
  if value >= int(math.Pow(2, float64(47))) {
    panic(fmt.Sprintf("number %d cannot be stored in 47 bits", value))
  }

  return Bit47{Value: value}
}

func (Bit47) Size() int {
  return 47
}

func (b Bit47) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit48 struct {
  Value int
}

func NewBit48(value int) Bit48 {
  if value >= int(math.Pow(2, float64(48))) {
    panic(fmt.Sprintf("number %d cannot be stored in 48 bits", value))
  }

  return Bit48{Value: value}
}

func (Bit48) Size() int {
  return 48
}

func (b Bit48) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit49 struct {
  Value int
}

func NewBit49(value int) Bit49 {
  if value >= int(math.Pow(2, float64(49))) {
    panic(fmt.Sprintf("number %d cannot be stored in 49 bits", value))
  }

  return Bit49{Value: value}
}

func (Bit49) Size() int {
  return 49
}

func (b Bit49) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit50 struct {
  Value int
}

func NewBit50(value int) Bit50 {
  if value >= int(math.Pow(2, float64(50))) {
    panic(fmt.Sprintf("number %d cannot be stored in 50 bits", value))
  }

  return Bit50{Value: value}
}

func (Bit50) Size() int {
  return 50
}

func (b Bit50) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit51 struct {
  Value int
}

func NewBit51(value int) Bit51 {
  if value >= int(math.Pow(2, float64(51))) {
    panic(fmt.Sprintf("number %d cannot be stored in 51 bits", value))
  }

  return Bit51{Value: value}
}

func (Bit51) Size() int {
  return 51
}

func (b Bit51) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit52 struct {
  Value int
}

func NewBit52(value int) Bit52 {
  if value >= int(math.Pow(2, float64(52))) {
    panic(fmt.Sprintf("number %d cannot be stored in 52 bits", value))
  }

  return Bit52{Value: value}
}

func (Bit52) Size() int {
  return 52
}

func (b Bit52) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit53 struct {
  Value int
}

func NewBit53(value int) Bit53 {
  if value >= int(math.Pow(2, float64(53))) {
    panic(fmt.Sprintf("number %d cannot be stored in 53 bits", value))
  }

  return Bit53{Value: value}
}

func (Bit53) Size() int {
  return 53
}

func (b Bit53) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit54 struct {
  Value int
}

func NewBit54(value int) Bit54 {
  if value >= int(math.Pow(2, float64(54))) {
    panic(fmt.Sprintf("number %d cannot be stored in 54 bits", value))
  }

  return Bit54{Value: value}
}

func (Bit54) Size() int {
  return 54
}

func (b Bit54) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit55 struct {
  Value int
}

func NewBit55(value int) Bit55 {
  if value >= int(math.Pow(2, float64(55))) {
    panic(fmt.Sprintf("number %d cannot be stored in 55 bits", value))
  }

  return Bit55{Value: value}
}

func (Bit55) Size() int {
  return 55
}

func (b Bit55) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit56 struct {
  Value int
}

func NewBit56(value int) Bit56 {
  if value >= int(math.Pow(2, float64(56))) {
    panic(fmt.Sprintf("number %d cannot be stored in 56 bits", value))
  }

  return Bit56{Value: value}
}

func (Bit56) Size() int {
  return 56
}

func (b Bit56) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit57 struct {
  Value int
}

func NewBit57(value int) Bit57 {
  if value >= int(math.Pow(2, float64(57))) {
    panic(fmt.Sprintf("number %d cannot be stored in 57 bits", value))
  }

  return Bit57{Value: value}
}

func (Bit57) Size() int {
  return 57
}

func (b Bit57) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit58 struct {
  Value int
}

func NewBit58(value int) Bit58 {
  if value >= int(math.Pow(2, float64(58))) {
    panic(fmt.Sprintf("number %d cannot be stored in 58 bits", value))
  }

  return Bit58{Value: value}
}

func (Bit58) Size() int {
  return 58
}

func (b Bit58) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit59 struct {
  Value int
}

func NewBit59(value int) Bit59 {
  if value >= int(math.Pow(2, float64(59))) {
    panic(fmt.Sprintf("number %d cannot be stored in 59 bits", value))
  }

  return Bit59{Value: value}
}

func (Bit59) Size() int {
  return 59
}

func (b Bit59) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit60 struct {
  Value int
}

func NewBit60(value int) Bit60 {
  if value >= int(math.Pow(2, float64(60))) {
    panic(fmt.Sprintf("number %d cannot be stored in 60 bits", value))
  }

  return Bit60{Value: value}
}

func (Bit60) Size() int {
  return 60
}

func (b Bit60) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit61 struct {
  Value int
}

func NewBit61(value int) Bit61 {
  if value >= int(math.Pow(2, float64(61))) {
    panic(fmt.Sprintf("number %d cannot be stored in 61 bits", value))
  }

  return Bit61{Value: value}
}

func (Bit61) Size() int {
  return 61
}

func (b Bit61) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit62 struct {
  Value int
}

func NewBit62(value int) Bit62 {
  if value >= int(math.Pow(2, float64(62))) {
    panic(fmt.Sprintf("number %d cannot be stored in 62 bits", value))
  }

  return Bit62{Value: value}
}

func (Bit62) Size() int {
  return 62
}

func (b Bit62) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit63 struct {
  Value int
}

func NewBit63(value int) Bit63 {
  if value >= int(math.Pow(2, float64(63))) {
    panic(fmt.Sprintf("number %d cannot be stored in 63 bits", value))
  }

  return Bit63{Value: value}
}

func (Bit63) Size() int {
  return 63
}

func (b Bit63) Serialize() uint64 {
  return uint64(b.Value)
}

type Bit64 struct {
  Value int
}

func NewBit64(value int) Bit64 {
  if value >= int(math.Pow(2, float64(64))) {
    panic(fmt.Sprintf("number %d cannot be stored in 64 bits", value))
  }

  return Bit64{Value: value}
}

func (Bit64) Size() int {
  return 64
}

func (b Bit64) Serialize() uint64 {
  return uint64(b.Value)
}