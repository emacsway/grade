package gradelogentry

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewReason(reason string) (Reason, error) {
	return Reason(reason), nil
}

type Reason string

func (r Reason) Export() string {
	return string(r)
}

func (r Reason) ExportTo(ex seedwork.ExporterSetter[string]) {
	ex.SetState(string(r))
}
