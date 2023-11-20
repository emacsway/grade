package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewCompetenceIdFaker() CompetenceIdFaker {
	return CompetenceIdFaker{
		TenantId:     tenant.TenantIdFakeValue,
		CompetenceId: uint(3),
	}
}

type CompetenceIdFaker struct {
	TenantId     uint
	CompetenceId uint
}

func (f CompetenceIdFaker) Create() (CompetenceId, error) {
	return NewCompetenceId(f.TenantId, f.CompetenceId)
}
