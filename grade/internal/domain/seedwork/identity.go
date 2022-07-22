package seedwork

import "github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"

func NewIdentity[T comparable](value T) (Identity[T, interfaces.Identity[T], interfaces.Exporter[T]], error) {
	return Identity[T, interfaces.Identity[T], interfaces.Exporter[T]]{value: value}, nil
}

// The way to fix issue of generics:
// https://issuemode.com/issues/golang/go/105227904

type Identity[T comparable, C interfaces.Identity[T], D interfaces.Exporter[T]] struct {
	value T
}

func (id Identity[T, C, D]) Equals(other C) bool {
	return id.value == other.Export()
}

func (id Identity[T, C, D]) Export() T {
	return id.value
}

func (id Identity[T, C, D]) ExportTo(ex D) {
	ex.SetState(id.value)
}
