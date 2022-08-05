package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

var MemberIdFakeValue = uuid.ParseSilent("cf9462cf-51d3-4c0a-b5ef-53b3dfccc7f6")

func NewTenantMemberIdFakeFactory() TenantMemberIdFakeFactory {
	return TenantMemberIdFakeFactory{
		TenantId: tenant.TenantIdFakeValue,
		MemberId: MemberIdFakeValue,
	}
}

type TenantMemberIdFakeFactory struct {
	TenantId uuid.Uuid
	MemberId uuid.Uuid
}

func (f *TenantMemberIdFakeFactory) NextTenantId() {
	f.TenantId = uuid.NewUuid()
}

func (f *TenantMemberIdFakeFactory) NextMemberId() {
	f.MemberId = uuid.NewUuid()
}

func (f TenantMemberIdFakeFactory) Create() (TenantMemberId, error) {
	return NewTenantMemberId(f.TenantId, f.MemberId)
}
