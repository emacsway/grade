package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/competence"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/faker"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
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

func WithTransientId() ArtifactFakerOption {
	return func(f *ArtifactFaker) {
		f.Id.ArtifactId = 0
	}
}

func WithRepository(repo ArtifactRepository) ArtifactFakerOption {
	return func(f *ArtifactFaker) {
		f.Repository = repo
	}
}

func WithMemberFaker(memberFaker *member.MemberFaker) ArtifactFakerOption {
	return func(f *ArtifactFaker) {
		f.MemberFaker = memberFaker
	}
}

func WithCompetenceFaker(competenceFaker *competence.CompetenceFaker) ArtifactFakerOption {
	return func(f *ArtifactFaker) {
		f.CompetenceFaker = competenceFaker
	}
}

func NewArtifactFaker(opts ...ArtifactFakerOption) *ArtifactFaker {
	f := &ArtifactFaker{
		Id:              values.NewArtifactIdFaker(),
		CompetenceIds:   []competenceVal.CompetenceIdFaker{competenceVal.NewCompetenceIdFaker()},
		AuthorIds:       []memberVal.MemberIdFaker{},
		OwnerId:         memberVal.NewMemberIdFaker(),
		MemberFaker:     member.NewMemberFaker(),
		CompetenceFaker: competence.NewCompetenceFaker(),
	}
	f.fake()
	repo := &ArtifactDummyRepository{
		IdFaker: &f.Id,
	}
	f.Repository = repo
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type ArtifactFaker struct {
	Id              values.ArtifactIdFaker
	Status          values.Status
	Name            string
	Description     string
	Url             string
	CompetenceIds   []competenceVal.CompetenceIdFaker
	AuthorIds       []memberVal.MemberIdFaker
	OwnerId         memberVal.MemberIdFaker
	CreatedAt       time.Time
	Repository      ArtifactRepository
	MemberFaker     *member.MemberFaker
	CompetenceFaker *competence.CompetenceFaker
	agg             *Artifact
}

func (f *ArtifactFaker) fake() {
	aFaker := faker.NewFaker()
	f.Status = values.Accepted
	f.Name = aFaker.Artifact()
	f.Description = aFaker.Sentences()
	f.Url = aFaker.Url()
	f.CreatedAt = time.Now().Truncate(time.Microsecond)
}

func (f *ArtifactFaker) Next() error {
	f.fake()
	f.AuthorIds = []memberVal.MemberIdFaker{}
	err := f.advanceId()
	if err != nil {
		return err
	}
	f.agg = nil
	return nil
}

func (f *ArtifactFaker) advanceId() error {
	var idExp values.ArtifactIdExporter
	tenantId, err := tenantVal.NewTenantId(f.Id.TenantId)
	if err != nil {
		return err
	}
	id, err := f.Repository.NextId(tenantId)
	if err != nil {
		return err
	}
	id.Export(&idExp)
	f.Id.ArtifactId = uint(idExp.ArtifactId)
	return nil
}

func (f *ArtifactFaker) AddAuthorId(authorId memberVal.MemberIdFaker) {
	f.AuthorIds = append(f.AuthorIds, authorId)
}

func (f *ArtifactFaker) AddCompetenceId(competenceId competenceVal.CompetenceIdFaker) {
	f.CompetenceIds = append(f.CompetenceIds, competenceId)
}

func (f *ArtifactFaker) Create() (*Artifact, error) {
	if f.agg != nil {
		return f.agg, nil
	}
	if f.Id.ArtifactId == 0 {
		err := f.advanceId()
		if err != nil {
			return nil, err
		}
	}
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
	var competenceIds []competenceVal.CompetenceId
	for i := range f.CompetenceIds {
		competenceId, err := f.CompetenceIds[i].Create()
		if err != nil {
			return nil, err
		}
		competenceIds = append(competenceIds, competenceId)
	}
	var authorIds []memberVal.MemberId
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
	agg, err := NewArtifact(
		id, f.Status, name, description, url,
		competenceIds, authorIds, owner, f.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(agg, aggregate.EventMeta{})
	if err != nil {
		return nil, err
	}
	f.agg = agg
	return agg, nil
}

// unidirectional flow of changes
func (f *ArtifactFaker) SetTenantId(val uint) {
	f.CompetenceFaker.SetTenantId(val)
}

func (f *ArtifactFaker) SetMemberId(val uint) {
	f.CompetenceFaker.SetMemberId(val)
}

func (f *ArtifactFaker) SetId(id memberVal.MemberIdFaker) {
	f.SetTenantId(id.TenantId)
	f.SetMemberId(id.MemberId)
}

func (f *ArtifactFaker) BuildDependencies() (err error) {
	err = f.CompetenceFaker.BuildDependencies()
	if err != nil {
		return err
	}
	_, err = f.CompetenceFaker.Create()
	if err != nil {
		return err
	}
	f.AddCompetenceId(f.CompetenceFaker.Id)
	f.Id.TenantId = f.CompetenceFaker.Id.TenantId
	f.OwnerId = f.CompetenceFaker.OwnerId
	*f.MemberFaker = *f.CompetenceFaker.MemberFaker
	return err
}

type ArtifactRepository interface {
	Insert(*Artifact, aggregate.EventMeta) error
	NextId(tenantVal.TenantId) (values.ArtifactId, error)
}

type ArtifactDummyRepository struct {
	IdFaker *values.ArtifactIdFaker
}

func (r ArtifactDummyRepository) Insert(agg *Artifact, eventMeta aggregate.EventMeta) error {
	return nil
}

func (r *ArtifactDummyRepository) NextId(tenantId tenantVal.TenantId) (values.ArtifactId, error) {
	var tenantIdExp exporters.UintExporter
	tenantId.Export(&tenantIdExp)
	r.IdFaker.TenantId = uint(tenantIdExp)
	r.IdFaker.ArtifactId += 1
	return r.IdFaker.Create()
}

type ArtifactDependencyFaker interface {
	MakeComptetnceFaker() *competence.CompetenceFaker
	MakeMemberFaker() *endorser.EndorserFaker
}
