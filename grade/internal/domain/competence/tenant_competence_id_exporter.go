package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantCompetenceIdExporter(tenantId, competenceId uuid.Uuid) TenantCompetenceIdExporter {
	return TenantCompetenceIdExporter{
		TenantId:     exporters.UuidExporter(tenantId),
		CompetenceId: exporters.UuidExporter(competenceId),
	}
}

type TenantCompetenceIdExporter struct {
	TenantId     exporters.UuidExporter
	CompetenceId exporters.UuidExporter
}

func (ex *TenantCompetenceIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantCompetenceIdExporter) SetCompetenceId(val CompetenceId) {
	val.Export(&ex.CompetenceId)
}
