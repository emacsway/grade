package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

type TenantMemberIdFakerOption func(*TenantMemberIdFaker)

func WithTenantId(tenantId uint) TenantMemberIdFakerOption {
	return func(f *TenantMemberIdFaker) {
		f.TenantId = tenantId
	}
}

func WithTransientId() TenantMemberIdFakerOption {
	return func(f *TenantMemberIdFaker) {
		f.MemberId = 0
	}
}

var MemberIdFakeValue = uint(3)

func NewTenantMemberIdFaker(opts ...TenantMemberIdFakerOption) TenantMemberIdFaker {
	f := TenantMemberIdFaker{
		TenantId: tenant.TenantIdFakeValue,
		MemberId: MemberIdFakeValue,
	}
	for _, opt := range opts {
		opt(&f)
	}
	return f
}

type TenantMemberIdFaker struct {
	TenantId uint
	MemberId uint
}

func (f *TenantMemberIdFaker) NextMemberId() {
	f.MemberId += 1
}

func (f TenantMemberIdFaker) Create() (TenantMemberId, error) {
	return NewTenantMemberId(f.TenantId, f.MemberId)
}
