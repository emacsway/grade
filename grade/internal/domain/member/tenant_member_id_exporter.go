package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantMemberIdExporter(tenantId, memberId uuid.Uuid) TenantMemberIdExporter {
	return TenantMemberIdExporter{
		TenantId: seedwork.UuidExporter(tenantId),
		MemberId: seedwork.UuidExporter(memberId),
	}
}

type TenantMemberIdExporter struct {
	TenantId seedwork.UuidExporter
	MemberId seedwork.UuidExporter
}

func (ex *TenantMemberIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantMemberIdExporter) SetMemberId(val MemberId) {
	val.Export(&ex.MemberId)
}
