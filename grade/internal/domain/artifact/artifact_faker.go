package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type ArtifactFakerOption func(*ArtifactFaker)

func WithTenantId(tenantId uint) ArtifactFakerOption {
	return func(f *ArtifactFaker) {
		f.Id.TenantId = tenantId
	}
}

func WithArtifactId(artifactId uint) ArtifactFakerOption {
	return func(f *ArtifactFaker) {
		f.Id.ArtifactId = artifactId
	}
}

func WithRepository(repo ArtifactRepository) ArtifactFakerOption {
	return func(f *ArtifactFaker) {
		f.Repository = repo
	}
}

func NewArtifactFaker(opts ...ArtifactFakerOption) *ArtifactFaker {
	f := &ArtifactFaker{
		Id:            values.NewTenantArtifactIdFaker(),
		Status:        values.Accepted,
		Name:          "Name1",
		Description:   "Description1",
		Url:           "https://github.com/emacsway/grade",
		CompetenceIds: []competence.TenantCompetenceIdFaker{competence.NewTenantCompetenceIdFaker()},
		AuthorIds:     []member.TenantMemberIdFaker{},
		OwnerId:       member.NewTenantMemberIdFaker(),
		CreatedAt:     time.Now().Truncate(time.Microsecond),
		Repository:    ArtifactDummyRepository{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type ArtifactFaker struct {
	Id            values.TenantArtifactIdFaker
	Status        values.Status
	Name          string
	Description   string
	Url           string
	CompetenceIds []competence.TenantCompetenceIdFaker
	AuthorIds     []member.TenantMemberIdFaker
	OwnerId       member.TenantMemberIdFaker
	CreatedAt     time.Time
	Repository    ArtifactRepository
}

func (f *ArtifactFaker) AddAuthorId(authorId member.TenantMemberIdFaker) error {
	// FIXME: return a error if the authorId already present in the list.
	f.AuthorIds = append(f.AuthorIds, authorId)
	return nil
}

func (f *ArtifactFaker) AddCompetenceId(competenceId competence.TenantCompetenceIdFaker) error {
	// FIXME: return a error if the authorId already present in the list.
	f.CompetenceIds = append(f.CompetenceIds, competenceId)
	return nil
}

func (f ArtifactFaker) Create() (*Artifact, error) {
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

func (f *ArtifactFaker) Next() error {
	f.Id.ArtifactId += 1
	return nil
}

type ArtifactRepository interface {
	Insert(*Artifact) error
}

type ArtifactDummyRepository struct{}

func (r ArtifactDummyRepository) Insert(agg *Artifact) error {
	return nil
}
