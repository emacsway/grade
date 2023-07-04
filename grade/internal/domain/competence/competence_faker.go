package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type CompetenceFakerOption func(*CompetenceFaker)

func WithTenantId(tenantId uint) CompetenceFakerOption {
	return func(f *CompetenceFaker) {
		f.Id.TenantId = tenantId
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

func NewCompetenceFaker(opts ...CompetenceFakerOption) *CompetenceFaker {
	f := &CompetenceFaker{
		Id:         values.NewTenantCompetenceIdFaker(),
		Name:       "Name1",
		OwnerId:    member.NewTenantMemberIdFaker(),
		CreatedAt:  time.Now().Truncate(time.Microsecond),
		Repository: CompetenceDummyRepository{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type CompetenceFaker struct {
	Id         values.TenantCompetenceIdFaker
	Name       string
	OwnerId    member.TenantMemberIdFaker
	CreatedAt  time.Time
	Repository CompetenceRepository
}

func (f *CompetenceFaker) Create() (*Competence, error) {
	var aggExp CompetenceExporter
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
	err = f.Repository.Save(agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id.CompetenceId = uint(aggExp.Id.CompetenceId)
	return agg, nil
}

type CompetenceRepository interface {
	Save(*Competence) error
}

type CompetenceDummyRepository struct{}

func (r CompetenceDummyRepository) Save(agg *Competence) error {
	return nil
}
