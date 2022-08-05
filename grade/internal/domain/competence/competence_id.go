package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewCompetenceId(value uuid.Uuid) (CompetenceId, error) {
	id, err := identity.NewUuidIdentity(value)
	if err != nil {
		return CompetenceId{}, err
	}
	return CompetenceId{id}, nil
}

type CompetenceId struct {
	identity.UuidIdentity
}
