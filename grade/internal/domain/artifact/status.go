package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

type Status uint8

func (s Status) Export(ex identity.ExporterSetter[uint8]) {
	ex.SetState(uint8(s))
}

const (
	Proposed = Status(0)
	Accepted = Status(1)
)
