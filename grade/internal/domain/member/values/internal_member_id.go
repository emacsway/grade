package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewInternalMemberId(value uint) (InternalMemberId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return InternalMemberId{}, err
	}
	return InternalMemberId{&id}, nil
}

func NewTransientInternalMemberId() InternalMemberId {
	return InternalMemberId{}
}

type InternalMemberId struct {
	*identity.IntIdentity
}
