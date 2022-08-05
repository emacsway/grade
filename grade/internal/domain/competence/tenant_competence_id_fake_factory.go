package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewTenantCompetenceIdFakeFactory() TenantCompetenceIdFakeFactory {
	return TenantCompetenceIdFakeFactory{
		TenantId:     uuid.ParseSilent("63e8d541-af30-4593-a8ac-761dc268926d"),
		CompetenceId: uuid.ParseSilent("cf9462cf-51d3-4c0a-b5ef-53b3dfccc7f6"),
	}
}

type TenantCompetenceIdFakeFactory struct {
	TenantId     uuid.Uuid
	CompetenceId uuid.Uuid
}

func (f TenantCompetenceIdFakeFactory) Create() (TenantCompetenceId, error) {
	return NewTenantCompetenceId(f.TenantId, f.CompetenceId)
}
