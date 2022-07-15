package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewRecognizerId(value uint64) RecognizerId {
	return RecognizerId{seedwork.NewIdentity(value)}
}

type RecognizerId struct {
	seedwork.Identity
}
