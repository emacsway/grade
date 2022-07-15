package seedwork

import "github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"

func NewIdentity(value uint64) Identity {
	return Identity{value: value}
}

type Identity struct {
	value uint64
}

func (id Identity) Equals(other interfaces.Identity[uint64]) bool {
	return id.value == other.GetValue()
}

func (id Identity) GetValue() uint64 {
	return id.value
}
