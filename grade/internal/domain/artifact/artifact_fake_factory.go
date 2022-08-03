package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/member"
)

func NewArtifactFakeFactory() ArtifactFakeFactory {
	idFactory := NewTenantArtifactIdFakeFactory()
	idFactory.ArtifactId = 20
	competenceIdsFactory := competence.NewTenantCompetenceIdFakeFactory()
	return ArtifactFakeFactory{
		Id:            idFactory,
		Status:        Accepted,
		Name:          "Name1",
		Description:   "Description1",
		Url:           "https://github.com/emacsway/grade",
		CompetenceIds: []competence.TenantCompetenceIdFakeFactory{competenceIdsFactory},
		AuthorIds:     []member.TenantMemberIdFakeFactory{},
		OwnerId:       member.NewTenantMemberIdFakeFactory(),
		CreatedAt:     time.Now(),
	}
}

type ArtifactFakeFactory struct {
	Id            TenantArtifactIdFakeFactory
	Status        Status
	Name          string
	Description   string
	Url           string
	CompetenceIds []competence.TenantCompetenceIdFakeFactory
	AuthorIds     []member.TenantMemberIdFakeFactory
	OwnerId       member.TenantMemberIdFakeFactory
	CreatedAt     time.Time
}

func (f *ArtifactFakeFactory) AddAuthorId(authorId member.TenantMemberIdFakeFactory) error {
	// FIXME: return a error if the authorId already present in the list.
	f.AuthorIds = append(f.AuthorIds, authorId)
	return nil
}

func (f *ArtifactFakeFactory) AddCompetenceId(competenceId competence.TenantCompetenceIdFakeFactory) error {
	// FIXME: return a error if the authorId already present in the list.
	f.CompetenceIds = append(f.CompetenceIds, competenceId)
	return nil
}

func (f ArtifactFakeFactory) Create() (*Artifact, error) {
	id, err := NewTenantArtifactId(f.Id.TenantId, f.Id.ArtifactId)
	if err != nil {
		return nil, err
	}
	name, err := NewName(f.Name)
	if err != nil {
		return nil, err
	}
	description, err := NewDescription(f.Description)
	if err != nil {
		return nil, err
	}
	url, err := NewUrl(f.Url)
	if err != nil {
		return nil, err
	}
	var competenceIds []competence.TenantCompetenceId
	for i := range f.CompetenceIds {
		competenceId, err := f.CompetenceIds[i].Create()
		if err != nil {
			return nil, err
		}
		competenceIds = append(competenceIds, competenceId)
	}
	var authorIds []member.TenantMemberId
	for i := range f.AuthorIds {
		authorId, err := f.AuthorIds[i].Create()
		if err != nil {
			return nil, err
		}
		authorIds = append(authorIds, authorId)
	}
	owner, err := f.OwnerId.Create()
	if err != nil {
		return nil, err
	}
	return NewArtifact(
		id, Accepted, name, description, url,
		competenceIds, authorIds, owner, f.CreatedAt,
	), nil
}
