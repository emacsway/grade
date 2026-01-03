package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewMemberIdExporter(tenantId, memberId uint) MemberIdExporter {
	return MemberIdExporter{
		TenantId: tenantId,
		MemberId: memberId,
	}
}

type MemberIdExporter struct {
	TenantId uint
	MemberId uint
}

func (ex *MemberIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(func(v uint) { ex.TenantId = v })
}

func (ex *MemberIdExporter) SetMemberId(val InternalMemberId) {
	val.Export(func(v uint) { ex.MemberId = v })
}
