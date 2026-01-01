package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func NewMemberIdExporter(tenantId, memberId uint) MemberIdExporter {
	return MemberIdExporter{
		TenantId: exporters.UintExporter(tenantId),
		MemberId: exporters.UintExporter(memberId),
	}
}

type MemberIdExporter struct {
	TenantId exporters.UintExporter
	MemberId exporters.UintExporter
}

func (ex *MemberIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *MemberIdExporter) SetMemberId(val InternalMemberId) {
	val.Export(&ex.MemberId)
}
