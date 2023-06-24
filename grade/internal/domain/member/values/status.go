package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type Status uint8

func (s Status) Export(ex exporters.ExporterSetter[uint8]) {
	ex.SetState(uint8(s))
}

const (
	Inactive = Status(0)
	Active   = Status(1)
)
