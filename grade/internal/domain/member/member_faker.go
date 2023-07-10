package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
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

func WithTenantFaker(tenantFaker *tenant.TenantFaker) MemberFakerOption {
	return func(f *MemberFaker) {
		// f.SetTenantFaker(tenantFaker)
	}
}

func NewMemberFaker(opts ...MemberFakerOption) *MemberFaker {
	f := &MemberFaker{
		Id:         values.NewTenantMemberIdFaker(),
		Repository: &MemberDummyRepository{},
	}
	f.fake()
	// f.SetTenantFaker(tenant.NewTenantFaker())
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type MemberFaker struct {
	Id        values.TenantMemberIdFaker
	Status    values.Status
	FullName  values.FullNameFaker
	CreatedAt time.Time
	// Repo and dependecies should be at Aggregate-level Faker, not at TenantMemberIdFaker
	Repository  MemberRepository
	TenantFaker *tenant.TenantFaker
	agg         *Member
}

func (f *MemberFaker) CreateDependencies() (err error) {
	_, err = f.TenantFaker.Create() // Use repo if it is needed to get an instance.
	if err != nil {
		return err
	}
	f.SetTenantFaker(f.TenantFaker)
	return err
}

func (f *MemberFaker) SetTenantFaker(tenantFaker *tenant.TenantFaker) {
	f.TenantFaker = tenantFaker
	f.Id.TenantId = f.TenantFaker.Id
}

func (f *MemberFaker) fake() {
	f.Status = values.Active
	f.FullName.Next()
	f.CreatedAt = time.Now().Truncate(time.Microsecond)
}

func (f *MemberFaker) Next() {
	f.fake()
	f.Id.MemberId += 1
	f.agg = nil
}

func (f *MemberFaker) Create() (*Member, error) {
	var aggExp MemberExporter
	if f.agg != nil {
		return f.agg, nil
	}

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
	f.agg = agg
	return agg, nil
}

type MemberRepository interface {
	Insert(*Member) error
}

type MemberDummyRepository struct{}

func (r *MemberDummyRepository) Insert(agg *Member) error {
	return nil
}
