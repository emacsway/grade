package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewTenantCompetenceIdExporter(tenantId, competenceId uint) TenantCompetenceIdExporter {
	return TenantCompetenceIdExporter{
		TenantId:     exporters.UintExporter(tenantId),
		CompetenceId: exporters.UintExporter(competenceId),
	}
}

type TenantCompetenceIdExporter struct {
	TenantId     exporters.UintExporter
	CompetenceId exporters.UintExporter
}

func (ex *TenantCompetenceIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantCompetenceIdExporter) SetCompetenceId(val CompetenceId) {
	val.Export(&ex.CompetenceId)
}