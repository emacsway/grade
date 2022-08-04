package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantCompetenceIdExporter(tenantId, competenceId uint64) TenantCompetenceIdExporter {
	return TenantCompetenceIdExporter{
		TenantId:     seedwork.Uint64Exporter(tenantId),
		CompetenceId: seedwork.Uint64Exporter(competenceId),
	}
}

type TenantCompetenceIdExporter struct {
	TenantId     seedwork.Uint64Exporter
	CompetenceId seedwork.Uint64Exporter
}

func (ex *TenantCompetenceIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantCompetenceIdExporter) SetCompetenceId(val CompetenceId) {
	val.Export(&ex.CompetenceId)
}
