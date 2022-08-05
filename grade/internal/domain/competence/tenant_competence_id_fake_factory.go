package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantCompetenceIdFakeFactory() TenantCompetenceIdFakeFactory {
	return TenantCompetenceIdFakeFactory{
		TenantId:     tenant.TenantIdFakeValue,
		CompetenceId: uuid.ParseSilent("cf9462cf-51d3-4c0a-b5ef-53b3dfccc7f6"),
	}
}

type TenantCompetenceIdFakeFactory struct {
	TenantId     uuid.Uuid
	CompetenceId uuid.Uuid
}

func (f *TenantCompetenceIdFakeFactory) NextTenantId() {
	f.TenantId = uuid.NewUuid()
}

func (f *TenantCompetenceIdFakeFactory) NextCompetenceId() {
	f.CompetenceId = uuid.NewUuid()
}

func (f TenantCompetenceIdFakeFactory) Create() (TenantCompetenceId, error) {
	return NewTenantCompetenceId(f.TenantId, f.CompetenceId)
}
