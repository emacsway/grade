package external

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewMemberId(value uint64) (MemberId, error) {
	id, err := seedwork.NewIdentity[uint64](value)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId{id}, nil
}

type MemberId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64], interfaces.Exporter[uint64]]
}
