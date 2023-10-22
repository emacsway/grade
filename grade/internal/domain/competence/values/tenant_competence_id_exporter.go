package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewTenantCompetenceIdExporter(tenantId, competenceInTenantId uint) TenantCompetenceIdExporter {
	return TenantCompetenceIdExporter{
		TenantId:             exporters.UintExporter(tenantId),
		CompetenceInTenantId: exporters.UintExporter(competenceInTenantId),
	}
}

type TenantCompetenceIdExporter struct {
	TenantId             exporters.UintExporter
	CompetenceInTenantId exporters.UintExporter
}

func (ex *TenantCompetenceIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantCompetenceIdExporter) SetCompetenceInTenantId(val CompetenceInTenantId) {
	val.Export(&ex.CompetenceInTenantId)
}
