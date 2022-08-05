package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantCompetenceIdExporter(tenantId, competenceId uuid.Uuid) TenantCompetenceIdExporter {
	return TenantCompetenceIdExporter{
		TenantId:     seedwork.UuidExporter(tenantId),
		CompetenceId: seedwork.UuidExporter(competenceId),
	}
}

type TenantCompetenceIdExporter struct {
	TenantId     seedwork.UuidExporter
	CompetenceId seedwork.UuidExporter
}

func (ex *TenantCompetenceIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantCompetenceIdExporter) SetCompetenceId(val CompetenceId) {
	val.Export(&ex.CompetenceId)
}
