package bitpack

type bits interface {
	Size() int
	Serialize() uint64
}
