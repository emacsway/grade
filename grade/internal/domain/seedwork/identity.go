package seedwork

import "github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"

func NewIdentity[T comparable](value T) Identity[T] {
	return Identity[T]{value: value}
}

type Identity[T comparable] struct {
	value T
}

func (id Identity[T]) Equals(other interfaces.Identity[T]) bool {
	return id.value == other.GetValue()
}

func (id Identity[T]) GetValue() T {
	return id.value
}
