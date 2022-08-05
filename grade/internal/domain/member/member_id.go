package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewMemberId(value uuid.Uuid) (MemberId, error) {
	id, err := identity.NewUuidIdentity(value)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId{id}, nil
}

type MemberId struct {
	identity.UuidIdentity
}
