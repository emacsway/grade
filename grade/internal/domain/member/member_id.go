package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewMemberId(value uint64) (MemberId, error) {
	id, err := seedwork.NewIdentity[uint64](value)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId{id}, nil
}

type MemberId struct {
	seedwork.Identity[uint64, seedwork.ExporterSetter[uint64]]
}
