package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewRecognizerId(value uint64) RecognizerId {
	return RecognizerId{value: value}
}

type RecognizerId struct {
	value uint64
}

func (id RecognizerId) Equals(other interfaces.Identity[uint64]) bool {
	return id.value == other.GetValue()
}

func (id RecognizerId) GetValue() uint64 {
	return id.value
}
