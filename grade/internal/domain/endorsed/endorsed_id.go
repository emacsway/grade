package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewEndorsedId(value uint64) EndorsedId {
	return EndorsedId{seedwork.NewIdentity[uint64](value)}
}

type EndorsedId struct {
	seedwork.Identity[uint64]
}
