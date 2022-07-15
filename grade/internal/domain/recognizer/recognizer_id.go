package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewRecognizerId(value uint64) RecognizerId {
	return RecognizerId{seedwork.NewIdentity[uint64](value)}
}

type RecognizerId struct {
	seedwork.Identity[uint64]
}
