package external

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewUserId(value uint64) UserId {
	return UserId{seedwork.NewIdentity[uint64](value)}
}

type UserId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64]]
}
