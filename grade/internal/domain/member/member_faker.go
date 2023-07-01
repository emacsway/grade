package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
)

type MemberFakerOption func(*MemberFaker)

func WithTenantId(tenantId uint) MemberFakerOption {
	return func(f *MemberFaker) {
		f.Id.TenantId = tenantId
	}
}

func WithTransientId() MemberFakerOption {
	return func(f *MemberFaker) {
		f.Id.MemberId = 0
	}
}

func WithRepository(repo MemberRepository) MemberFakerOption {
	return func(f *MemberFaker) {
		f.Repository = repo
	}
}

func NewMemberFaker(opts ...MemberFakerOption) *MemberFaker {
	f := &MemberFaker{
		Id:        values.NewTenantMemberIdFaker(),
		Status:    values.Active,
		FullName:  values.NewFullNameFaker(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
		// Repo and dependecies should be at Aggregate-level Faker, not at TenantMemberIdFaker
		Repository: MemberDummyRepository{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type MemberFaker struct {
	Id         values.TenantMemberIdFaker
	Status     values.Status
	FullName   values.FullNameFaker
	CreatedAt  time.Time
	Repository MemberRepository
}

func (f *MemberFaker) Create() (*Member, error) {
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
	)
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
