package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
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
		f.TenantFaker = tenantFaker
	}
}

func NewMemberFaker(opts ...MemberFakerOption) *MemberFaker {
	f := &MemberFaker{
		Id:          values.NewMemberIdFaker(),
		Repository:  &MemberDummyRepository{},
		TenantFaker: tenant.NewTenantFaker(),
	}
	f.fake()
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type MemberFaker struct {
	Id        values.MemberIdFaker
	Status    values.Status
	FullName  values.FullNameFaker
	CreatedAt time.Time
	// Repo and dependecies should be at Aggregate-level Faker, not at MemberIdFaker
	Repository  MemberRepository
	TenantFaker *tenant.TenantFaker
	agg         *Member
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

func (f *MemberFaker) Create(s session.Session) (*Member, error) {
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
	err = f.Repository.Insert(s, agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id.MemberId = uint(aggExp.Id.MemberId)
	f.agg = agg
	return agg, nil
}

// unidirectional flow of changes
func (f *MemberFaker) SetTenantId(val uint) {
	f.TenantFaker.SetId(val)
}

func (f *MemberFaker) SetMemberId(val uint) {
	f.Id.MemberId = val
}

func (f *MemberFaker) SetId(id values.MemberIdFaker) {
	f.SetTenantId(id.TenantId)
	f.SetMemberId(id.MemberId)
}

func (f *MemberFaker) BuildDependencies(s session.Session) (err error) {
	_, err = f.TenantFaker.Create(s) // Use repo if it is needed to get the instance.
	if err != nil {
		return err
	}
	f.Id.TenantId = f.TenantFaker.Id
	return err
}

type MemberRepository interface {
	Insert(session.Session, *Member) error
}

type MemberDummyRepository struct{}

func (r *MemberDummyRepository) Insert(s session.Session, agg *Member) error {
	return nil
}
