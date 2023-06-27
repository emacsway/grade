package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

type TenantMemberIdFakeFactoryOption func(*TenantMemberIdFakeFactory)

func WithTenantId(tenantId uint) TenantMemberIdFakeFactoryOption {
	return func(f *TenantMemberIdFakeFactory) {
		f.TenantId = tenantId
	}
}

func WithTransientId() TenantMemberIdFakeFactoryOption {
	return func(f *TenantMemberIdFakeFactory) {
		f.MemberId = 0
	}
}

var MemberIdFakeValue = uint(3)

func NewTenantMemberIdFakeFactory(opts ...TenantMemberIdFakeFactoryOption) TenantMemberIdFakeFactory {
	f := TenantMemberIdFakeFactory{
		TenantId: tenant.TenantIdFakeValue,
		MemberId: MemberIdFakeValue,
	}
	for _, opt := range opts {
		opt(&f)
	}
	return f
}

type TenantMemberIdFakeFactory struct {
	TenantId uint
	MemberId uint
}

func (f *TenantMemberIdFakeFactory) NextMemberId() {
	f.MemberId += 1
}

func (f TenantMemberIdFakeFactory) Create() (TenantMemberId, error) {
	return NewTenantMemberId(f.TenantId, f.MemberId)
}
