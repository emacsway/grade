package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member"
)

func NewCompetenceFakeFactory() CompetenceFakeFactory {
	return CompetenceFakeFactory{
		Id:        NewTenantCompetenceIdFakeFactory(),
		Name:      "Name1",
		OwnerId:   member.NewTenantMemberIdFakeFactory(),
		CreatedAt: time.Now(),
	}
}

type CompetenceFakeFactory struct {
	Id        TenantCompetenceIdFakeFactory
	Name      string
	OwnerId   member.TenantMemberIdFakeFactory
	CreatedAt time.Time
}

func (f CompetenceFakeFactory) Create() (*Competence, error) {
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	name, err := NewName(f.Name)
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
