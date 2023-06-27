package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func NewCompetenceFakeFactory() *CompetenceFakeFactory {
	return &CompetenceFakeFactory{
		Id:        values.NewTenantCompetenceIdFakeFactory(),
		Name:      "Name1",
		OwnerId:   member.NewTenantMemberIdFakeFactory(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
}

type CompetenceFakeFactory struct {
	Id        values.TenantCompetenceIdFakeFactory
	Name      string
	OwnerId   member.TenantMemberIdFakeFactory
	CreatedAt time.Time
}

func (f CompetenceFakeFactory) Create() (*Competence, error) {
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
