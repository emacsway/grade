package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewCompetenceIdExporter(tenantId, competenceId uint) CompetenceIdExporter {
	return CompetenceIdExporter{
		TenantId:     exporters.UintExporter(tenantId),
		CompetenceId: exporters.UintExporter(competenceId),
	}
}

type CompetenceIdExporter struct {
	TenantId     exporters.UintExporter
	CompetenceId exporters.UintExporter
}

func (ex *CompetenceIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *CompetenceIdExporter) SetCompetenceId(val InternalCompetenceId) {
	val.Export(&ex.CompetenceId)
}
