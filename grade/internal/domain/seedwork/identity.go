package seedwork

import (
	"fmt"
)

type ExporterSetter[T any] interface {
	SetState(T)
}

type ExportableTo[T any] interface {
	Export(ExporterSetter[T])
}

// alternative approach:

type Accessable[T any] interface {
	Value() T
}

type Equaler interface {
	Equal(Equaler) bool
}

func NewUint64Identity(value uint64) (Uint64Identity, error) {
	return Uint64Identity{value: value}, nil
}

type Uint64Identity struct {
	value uint64
}

func (id Uint64Identity) Equal(other Equaler) bool {
	if id.value == 0 {
		return false // Aggregate is not saved.
	}
	exportableOther := other.(Accessable[uint64])
	return id.value == exportableOther.Value()
}

func (id Uint64Identity) Export(ex ExporterSetter[uint64]) {
	ex.SetState(id.value)
}

func (id Uint64Identity) Value() uint64 {
	return id.value
}

func (id Uint64Identity) String() string {
	return fmt.Sprintf("%d", id.value)
}
