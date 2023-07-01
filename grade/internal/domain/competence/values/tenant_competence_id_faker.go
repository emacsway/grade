package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewTenantCompetenceIdFaker() TenantCompetenceIdFaker {
	return TenantCompetenceIdFaker{
		TenantId:     tenant.TenantIdFakeValue,
		CompetenceId: uint(3),
	}
}

type TenantCompetenceIdFaker struct {
	TenantId     uint
	CompetenceId uint
}

func (f TenantCompetenceIdFaker) Create() (TenantCompetenceId, error) {
	return NewTenantCompetenceId(f.TenantId, f.CompetenceId)
}