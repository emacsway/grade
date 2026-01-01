package values

import (
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func NewStatus(val uint8) (Status, error) {
	return Status(val), nil
}

type Status uint8

func (s Status) Export(ex exporters.ExporterSetter[uint8]) {
	ex.SetState(uint8(s))
}

const (
	Proposed = Status(0)
	Accepted = Status(1)
)
