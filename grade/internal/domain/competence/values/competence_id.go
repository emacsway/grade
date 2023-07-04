package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewCompetenceId(value uint) (CompetenceId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return CompetenceId{}, err
	}
	return CompetenceId{&id}, nil
}

type CompetenceId struct {
	*identity.IntIdentity
}
