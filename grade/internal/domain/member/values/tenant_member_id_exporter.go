package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantMemberIdExporter(tenantId, memberId uint) TenantMemberIdExporter {
	return TenantMemberIdExporter{
		TenantId: exporters.UintExporter(tenantId),
		MemberId: exporters.UintExporter(memberId),
	}
}

type TenantMemberIdExporter struct {
	TenantId exporters.UintExporter
	MemberId exporters.UintExporter
}

func (ex *TenantMemberIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantMemberIdExporter) SetMemberId(val MemberId) {
	val.Export(&ex.MemberId)
}
