package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantCompetenceIdFakeFactory() TenantCompetenceIdFakeFactory {
	return TenantCompetenceIdFakeFactory{
		TenantId:     tenant.TenantIdFakeValue,
		CompetenceId: uint(3),
	}
}

type TenantCompetenceIdFakeFactory struct {
	TenantId     uint
	CompetenceId uint
}

func (f *TenantCompetenceIdFakeFactory) NextTenantId() {
	f.TenantId += 1
}

func (f *TenantCompetenceIdFakeFactory) NextCompetenceId() {
	f.CompetenceId += 1
}

func (f TenantCompetenceIdFakeFactory) Create() (TenantCompetenceId, error) {
	return NewTenantCompetenceId(f.TenantId, f.CompetenceId)
}
