package bitpack

type Arbitrary interface {
	Size() int
	Serialize() uint64
}
