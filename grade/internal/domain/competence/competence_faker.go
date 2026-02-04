package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/faker"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

type CompetenceFakerOption func(*CompetenceFaker)

func WithTenantId(tenantId uint) CompetenceFakerOption {
	return func(f *CompetenceFaker) {
		f.Id.TenantId = tenantId
		f.OwnerId.TenantId = tenantId
	}
}

func WithTransientId() CompetenceFakerOption {
	return func(f *CompetenceFaker) {
		f.Id.CompetenceId = 0
	}
}

func WithRepository(repo CompetenceRepository) CompetenceFakerOption {
	return func(f *CompetenceFaker) {
		f.Repository = repo
	}
}

func WithMemberFaker(memberFaker *member.MemberFaker) CompetenceFakerOption {
	return func(f *CompetenceFaker) {
		f.MemberFaker = memberFaker
	}
}

func NewCompetenceFaker(opts ...CompetenceFakerOption) *CompetenceFaker {
	f := &CompetenceFaker{
		Id:          values.NewCompetenceIdFaker(),
		OwnerId:     memberVal.NewMemberIdFaker(),
		Repository:  CompetenceDummyRepository{},
		MemberFaker: member.NewMemberFaker(),
	}
	f.fake()
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type CompetenceFaker struct {
	Id          values.CompetenceIdFaker
	Name        string
	OwnerId     memberVal.MemberIdFaker
	CreatedAt   time.Time
	Repository  CompetenceRepository
	MemberFaker *member.MemberFaker
	agg         *Competence
}

func (f *CompetenceFaker) fake() {
	aFaker := faker.NewFaker()
	f.Name = aFaker.Competence()
	f.CreatedAt = time.Now().Truncate(time.Microsecond)
}

func (f *CompetenceFaker) Next() error {
	f.fake()
	f.Id.CompetenceId += 1
	f.agg = nil
	return nil
}

func (f *CompetenceFaker) Create(s session.Session) (*Competence, error) {
	var aggExp CompetenceExporter
	if f.agg != nil {
		return f.agg, nil
	}

	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	name, err := values.NewName(f.Name)
	if err != nil {
		return nil, err
	}
	owner, err := f.OwnerId.Create()
	if err != nil {
		return nil, err
	}
	agg, err := NewCompetence(
		id, name, owner, f.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(s, agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id.CompetenceId = uint(aggExp.Id.CompetenceId)
	f.agg = agg
	return agg, nil
}

// unidirectional flow of changes
func (f *CompetenceFaker) SetTenantId(val uint) {
	f.MemberFaker.SetTenantId(val)
}

func (f *CompetenceFaker) SetMemberId(val uint) {
	f.MemberFaker.SetMemberId(val)
}

func (f *CompetenceFaker) SetId(id memberVal.MemberIdFaker) {
	f.SetTenantId(id.TenantId)
	f.SetMemberId(id.MemberId)
}

func (f *CompetenceFaker) BuildDependencies(s session.Session) (err error) {
	err = f.MemberFaker.BuildDependencies(s)
	if err != nil {
		return err
	}
	_, err = f.MemberFaker.Create(s)
	if err != nil {
		return err
	}
	f.Id.TenantId = f.MemberFaker.Id.TenantId
	f.OwnerId = f.MemberFaker.Id
	return err
}

type CompetenceRepository interface {
	Insert(session.Session, *Competence) error
}

type CompetenceDummyRepository struct{}

func (r CompetenceDummyRepository) Insert(s session.Session, agg *Competence) error {
	return nil
}
