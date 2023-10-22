package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewCompetenceInTenantId(value uint) (CompetenceInTenantId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return CompetenceInTenantId{}, err
	}
	return CompetenceInTenantId{&id}, nil
}

type CompetenceInTenantId struct {
	*identity.IntIdentity
}
