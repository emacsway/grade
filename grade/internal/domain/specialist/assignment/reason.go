package assignment

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewReason(reason string) (Reason, error) {
	return Reason(reason), nil
}

type Reason string

func (r Reason) Export(ex identity.ExporterSetter[string]) {
	ex.SetState(string(r))
}
