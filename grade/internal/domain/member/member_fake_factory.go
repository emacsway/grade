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

func WithRepository(repo MemberRepository) MemberFakeFactoryOption {
	return func(f *MemberFakeFactory) {
		f.Repository = repo
	}
}

func NewMemberFakeFactory(opts ...MemberFakeFactoryOption) *MemberFakeFactory {
	f := &MemberFakeFactory{
		Id:        values.NewTenantMemberIdFakeFactory(),
		Status:    values.Active,
		FullName:  values.NewFullNameFakeFactory(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
		// Repo and dependecies should be at Aggregate-level FakeFactory, not at TenantMemberIdFakeFactory
		Repository: MemberDummyRepository{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type MemberFakeFactory struct {
	Id         values.TenantMemberIdFakeFactory
	Status     values.Status
	FullName   values.FullNameFakeFactory
	CreatedAt  time.Time
	Repository MemberRepository
}

func (f MemberFakeFactory) Create() (*Member, error) {
	var aggExp MemberExporter
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	fullName, err := f.FullName.Create()
	if err != nil {
		return nil, err
	}
	agg, err := NewMember(
		id, f.Status, fullName, f.CreatedAt,
	), nil
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id.MemberId = uint(aggExp.Id.MemberId)
	return agg, nil
}

type MemberRepository interface {
	Insert(*Member) error
}

type MemberDummyRepository struct{}

func (r MemberDummyRepository) Insert(agg *Member) error {
	return nil
}
