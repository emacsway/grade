package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewMemberId(value int) MemberId {
	return MemberId{value: value}
}

type MemberId struct {
	value int
}

func (memberId MemberId) Equals(other interfaces.Identity[int]) bool {
	return memberId.value == other.GetValue()
}

func (memberId MemberId) GetValue() int {
	return memberId.value
}
