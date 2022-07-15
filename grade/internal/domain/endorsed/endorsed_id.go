package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewEndorsedId(value uint64) EndorsedId {
	return EndorsedId{value: value}
}

type EndorsedId struct {
	value uint64
}

func (id EndorsedId) Equals(other interfaces.Identity[uint64]) bool {
	return id.value == other.GetValue()
}

func (id EndorsedId) GetValue() uint64 {
	return id.value
}
