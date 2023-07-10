package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/faker"
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
		// TODO: f.SetMemberFaker(memberFaker)
	}
}

func NewCompetenceFaker(opts ...CompetenceFakerOption) *CompetenceFaker {
	f := &CompetenceFaker{
		Id:         values.NewTenantCompetenceIdFaker(),
		OwnerId:    memberVal.NewTenantMemberIdFaker(),
		Repository: CompetenceDummyRepository{},
	}
	f.fake()
	// TODO: f.SetMemberFaker(member.NewMemberFaker())
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type CompetenceFaker struct {
	Id          values.TenantCompetenceIdFaker
	Name        string
	OwnerId     memberVal.TenantMemberIdFaker
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

func (f *CompetenceFaker) Next() {
	f.fake()
	f.Id.CompetenceId += 1
	f.agg = nil
}

func (f *CompetenceFaker) Create() (*Competence, error) {
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
	err = f.Repository.Insert(agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id.CompetenceId = uint(aggExp.Id.CompetenceId)
	f.agg = agg
	return agg, nil
}

func (f *CompetenceFaker) BuildDependencies() (err error) {
	err = f.MemberFaker.BuildDependencies()
	if err != nil {
		return err
	}
	_, err = f.MemberFaker.Create()
	if err != nil {
		return err
	}
	f.SetMemberFaker(f.MemberFaker)
	return err
}

func (f *CompetenceFaker) SetMemberFaker(memberFaker *member.MemberFaker) {
	f.MemberFaker = memberFaker
	f.OwnerId = f.MemberFaker.Id
}

type CompetenceRepository interface {
	Insert(*Competence) error
}

type CompetenceDummyRepository struct{}

func (r CompetenceDummyRepository) Insert(agg *Competence) error {
	return nil
}
