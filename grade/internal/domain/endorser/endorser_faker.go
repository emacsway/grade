package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
)

var EndorserMemberIdFakeValue = uint(1004)

type EndorserFakerOption func(*EndorserFaker)

func WithTenantId(tenantId uint) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.Id.TenantId = tenantId
	}
}

func WithMemberId(memberId uint) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.Id.MemberId = memberId
	}
}

func WithRepository(repo EndorserRepository) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.Repository = repo
	}
}

func WithMemberFaker(memberFaker *member.MemberFaker) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.Dependency.MemberFaker = memberFaker
	}
}

func NewEndorserFaker(opts ...EndorserFakerOption) *EndorserFaker {
	idFactory := memberVal.NewTenantMemberIdFaker()
	idFactory.MemberId = EndorserMemberIdFakeValue
	f := &EndorserFaker{
		Id:        idFactory,
		Grade:     1,
		CreatedAt: time.Now().Truncate(time.Microsecond),
		Dependency: &Dependency{
			MemberFaker: member.NewMemberFaker(),
		},
	}
	repo := &EndorserDummyRepository{
		Faker: f,
	}
	f.Repository = repo
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type EndorserFaker struct {
	Id         memberVal.TenantMemberIdFaker
	Grade      uint8
	CreatedAt  time.Time
	Repository EndorserRepository
	Dependency *Dependency
}

func (f *EndorserFaker) CreateDependencies() error {
	return f.Dependency.Create(f)
}

func (f EndorserFaker) Create() (*Endorser, error) {
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	agg, err := NewEndorser(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	g, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return nil, err
	}
	err = agg.SetGrade(g)
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(agg)
	if err != nil {
		return nil, err
	}
	return agg, nil
}

type EndorserRepository interface {
	Insert(*Endorser) error
}

type EndorserDummyRepository struct {
	Faker *EndorserFaker
}

func (r EndorserDummyRepository) Insert(agg *Endorser) error {
	return nil
}

type Dependency struct {
	MemberFaker *member.MemberFaker
}

func (d *Dependency) Create(f *EndorserFaker) (err error) {
	d.MemberFaker.CreateDependencies()
	_, err = d.MemberFaker.Create()
	f.Id = d.MemberFaker.Id
	return err
}
