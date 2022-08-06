package assignment

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func NewReason(reason string) (Reason, error) {
	return Reason(reason), nil
}

type Reason string

func (r Reason) Export(ex exporters.ExporterSetter[string]) {
	ex.SetState(string(r))
}
