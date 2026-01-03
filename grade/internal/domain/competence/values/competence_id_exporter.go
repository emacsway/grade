package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewCompetenceIdExporter(tenantId, competenceId uint) CompetenceIdExporter {
	return CompetenceIdExporter{
		TenantId:     tenantId,
		CompetenceId: competenceId,
	}
}

type CompetenceIdExporter struct {
	TenantId     uint
	CompetenceId uint
}

func (ex *CompetenceIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(func(v uint) { ex.TenantId = v })
}

func (ex *CompetenceIdExporter) SetCompetenceId(val InternalCompetenceId) {
	val.Export(func(v uint) { ex.CompetenceId = v })
}
