package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

type MemberIdFakerOption func(*MemberIdFaker)

func WithTenantId(tenantId uint) MemberIdFakerOption {
	return func(f *MemberIdFaker) {
		f.TenantId = tenantId
	}
}

func WithTransientId() MemberIdFakerOption {
	return func(f *MemberIdFaker) {
		f.MemberId = 0
	}
}

var MemberIdFakeValue = uint(3)

func NewMemberIdFaker(opts ...MemberIdFakerOption) MemberIdFaker {
	f := MemberIdFaker{
		TenantId: tenant.TenantIdFakeValue,
		MemberId: MemberIdFakeValue,
	}
	for _, opt := range opts {
		opt(&f)
	}
	return f
}

type MemberIdFaker struct {
	TenantId uint
	MemberId uint
}

func (f *MemberIdFaker) NextMemberId() {
	f.MemberId += 1
}

func (f MemberIdFaker) Create() (MemberId, error) {
	return NewMemberId(f.TenantId, f.MemberId)
}
