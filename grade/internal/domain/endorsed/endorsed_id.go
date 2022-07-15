package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewEndorsedId(value uint64) EndorsedId {
	return EndorsedId{seedwork.NewIdentity(value)}
}

type EndorsedId struct {
	seedwork.Identity
}
