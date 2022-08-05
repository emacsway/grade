package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewMemberId(value uuid.Uuid) (MemberId, error) {
	id, err := seedwork.NewUuidIdentity(value)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId{id}, nil
}

type MemberId struct {
	seedwork.UuidIdentity
}
