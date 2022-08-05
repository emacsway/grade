package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantMemberIdExporter(tenantId, memberId uuid.Uuid) TenantMemberIdExporter {
	return TenantMemberIdExporter{
		TenantId: exporters.UuidExporter(tenantId),
		MemberId: exporters.UuidExporter(memberId),
	}
}

type TenantMemberIdExporter struct {
	TenantId exporters.UuidExporter
	MemberId exporters.UuidExporter
}

func (ex *TenantMemberIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantMemberIdExporter) SetMemberId(val MemberId) {
	val.Export(&ex.MemberId)
}
