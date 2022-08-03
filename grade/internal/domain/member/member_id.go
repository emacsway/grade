package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
)

func NewMemberId(value uint64) (MemberId, error) {
	id, err := seedwork.NewUint64Identity(value)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId{id}, nil
}

type MemberId struct {
	seedwork.Uint64Identity
}
