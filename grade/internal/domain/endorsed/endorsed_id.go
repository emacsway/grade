package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewEndorsedId(value uint64) EndorsedId {
	return EndorsedId{seedwork.NewIdentity[uint64](value)}
}

type EndorsedId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64]]
}
