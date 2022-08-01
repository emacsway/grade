package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/tenant"
)

func NewTenantMemberIdExporter(tenantId uint64, memberId uint64) TenantMemberIdExporter {
	return TenantMemberIdExporter{
		TenantId: seedwork.Uint64Exporter(tenantId),
		MemberId: seedwork.Uint64Exporter(memberId),
	}
}

type TenantMemberIdExporter struct {
	TenantId seedwork.Uint64Exporter
	MemberId seedwork.Uint64Exporter
}

func (ex *TenantMemberIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantMemberIdExporter) SetMemberId(val MemberId) {
	val.Export(&ex.MemberId)
}
