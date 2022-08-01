package assignment

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewReason(reason string) (Reason, error) {
	return Reason(reason), nil
}

type Reason string

func (r Reason) Export(ex seedwork.ExporterSetter[string]) {
	ex.SetState(string(r))
}
