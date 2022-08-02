package artifact

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/competence"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

func NewArtifactFakeFactory() ArtifactFakeFactory {
	idFactory := NewTenantArtifactIdFakeFactory()
	idFactory.ArtifactId = 20
	return ArtifactFakeFactory{
		Id:            idFactory,
		Status:        Accepted,
		Name:          "Name1",
		Description:   "Description1",
		Url:           "https://github.com/emacsway/qualifying-grade",
		CompetenceIds: []uint64{},
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
	CompetenceIds []uint64
	AuthorIds     []member.TenantMemberIdFakeFactory
	OwnerId       member.TenantMemberIdFakeFactory
	CreatedAt     time.Time
}

func (f *ArtifactFakeFactory) AddAuthorId(authorId member.TenantMemberIdFakeFactory) error {
	// FIXME: return a error if the authorId already present in the list.
	f.AuthorIds = append(f.AuthorIds, authorId)
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
		[]competence.CompetenceId{}, authorIds, owner, f.CreatedAt,
	), nil
}
