package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
)

type MemberFakeFactoryOption func(*MemberFakeFactory)

func WithTenantId(tenantId uint) MemberFakeFactoryOption {
	return func(f *MemberFakeFactory) {
		f.Id.TenantId = tenantId
	}
}

func WithTransientId() MemberFakeFactoryOption {
	return func(f *MemberFakeFactory) {
		f.Id.MemberId = 0
	}
}

func NewMemberFakeFactory(opts ...MemberFakeFactoryOption) MemberFakeFactory {
	f := MemberFakeFactory{
		Id:        values.NewTenantMemberIdFakeFactory(),
		Status:    values.Active,
		FullName:  values.NewFullNameFakeFactory(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
	for _, opt := range opts {
		opt(&f)
	}
	return f
}

type MemberFakeFactory struct {
	Id        values.TenantMemberIdFakeFactory
	Status    values.Status
	FullName  values.FullNameFakeFactory
	CreatedAt time.Time
}

func (f MemberFakeFactory) Create() (*Member, error) {
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	fullName, err := f.FullName.Create()
	if err != nil {
		return nil, err
	}
	return NewMember(
		id, f.Status, fullName, f.CreatedAt,
	), nil
}
