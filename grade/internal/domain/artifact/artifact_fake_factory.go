package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func NewArtifactFakeFactory() ArtifactFakeFactory {
	return ArtifactFakeFactory{
		Id:            values.NewTenantArtifactIdFakeFactory(),
		Status:        values.Accepted,
		Name:          "Name1",
		Description:   "Description1",
		Url:           "https://github.com/emacsway/grade",
		CompetenceIds: []competence.TenantCompetenceIdFakeFactory{competence.NewTenantCompetenceIdFakeFactory()},
		AuthorIds:     []member.TenantMemberIdFakeFactory{},
		OwnerId:       member.NewTenantMemberIdFakeFactory(),
		CreatedAt:     time.Now().Truncate(time.Microsecond),
	}
}

type ArtifactFakeFactory struct {
	Id            values.TenantArtifactIdFakeFactory
	Status        values.Status
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
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	name, err := values.NewName(f.Name)
	if err != nil {
		return nil, err
	}
	description, err := values.NewDescription(f.Description)
	if err != nil {
		return nil, err
	}
	url, err := values.NewUrl(f.Url)
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
		id, f.Status, name, description, url,
		competenceIds, authorIds, owner, f.CreatedAt,
	), nil
}
