package values

import (
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/identity"
)

func NewInternalCompetenceId(value uint) (InternalCompetenceId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return InternalCompetenceId{}, err
	}
	return InternalCompetenceId{&id}, nil
}

type InternalCompetenceId struct {
	*identity.IntIdentity
}
