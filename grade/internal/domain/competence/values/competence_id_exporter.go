package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
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
