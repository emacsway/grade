package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewMemberId(value uint) (MemberId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId{id}, nil
}

func NewTransientMemberId() MemberId {
	return MemberId{}
}

type MemberId struct {
	identity.IntIdentity
}
