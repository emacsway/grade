package values

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

var MemberIdFakeValue = uint(3)

func NewTenantMemberIdFakeFactory() TenantMemberIdFakeFactory {
	return TenantMemberIdFakeFactory{
		TenantId: tenant.TenantIdFakeValue,
		MemberId: MemberIdFakeValue,
	}
}

type TenantMemberIdFakeFactory struct {
	TenantId uint
	MemberId uint
}

func (f *TenantMemberIdFakeFactory) NextTenantId() {
	f.TenantId += 1
}

func (f *TenantMemberIdFakeFactory) NextMemberId() {
	f.MemberId += 1
}

func (f TenantMemberIdFakeFactory) Create() (TenantMemberId, error) {
	return NewTenantMemberId(f.TenantId, f.MemberId)
}
