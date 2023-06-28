package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func NewCompetenceFaker() *CompetenceFaker {
	return &CompetenceFaker{
		Id:        values.NewTenantCompetenceIdFaker(),
		Name:      "Name1",
		OwnerId:   member.NewTenantMemberIdFaker(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
}

type CompetenceFaker struct {
	Id        values.TenantCompetenceIdFaker
	Name      string
	OwnerId   member.TenantMemberIdFaker
	CreatedAt time.Time
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
