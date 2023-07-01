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

func NewCompetenceFaker() *CompetenceFaker {
	return &CompetenceFaker{
		Id:         values.NewTenantCompetenceIdFaker(),
		Name:       "Name1",
		OwnerId:    member.NewTenantMemberIdFaker(),
		CreatedAt:  time.Now().Truncate(time.Microsecond),
		Repository: CompetenceDummyRepository{},
	}
}

type CompetenceFaker struct {
	Id         values.TenantCompetenceIdFaker
	Name       string
	OwnerId    member.TenantMemberIdFaker
	CreatedAt  time.Time
	Repository CompetenceRepository
}

func (f CompetenceFaker) Create() (*Competence, error) {
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
	return NewCompetence(
		id, name, owner, f.CreatedAt,
	)
}

type CompetenceRepository interface {
	Insert(*Competence) error
}

type CompetenceDummyRepository struct{}

func (r CompetenceDummyRepository) Insert(agg *Competence) error {
	return nil
}
