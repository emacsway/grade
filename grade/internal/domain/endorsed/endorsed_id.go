package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewEndorsedId(value int) EndorsedId {
	return EndorsedId{value: value}
}

type EndorsedId struct {
	value int
}

func (id EndorsedId) Equals(other interfaces.Identity[int]) bool {
	return id.value == other.GetValue()
}

func (id EndorsedId) GetValue() int {
	return id.value
}
