package events

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

func NewArtifactProposed(
	id values.TenantArtifactId,
	status values.Status,
	name values.Name,
	description values.Description,
	url values.Url,
	competenceIds []competence.TenantCompetenceId,
	authorIds []member.TenantMemberId,
	ownerId member.TenantMemberId,
	createdAt time.Time,
) *ArtifactProposed {
	return &ArtifactProposed{
		id:            id,
		status:        status,
		name:          name,
		description:   description,
		url:           url,
		competenceIds: competenceIds,
		authorIds:     authorIds,
		ownerId:       ownerId,
		createdAt:     createdAt,
	}
}

type ArtifactProposed struct {
	id               values.TenantArtifactId
	status           values.Status
	name             values.Name
	description      values.Description
	url              values.Url
	competenceIds    []competence.TenantCompetenceId
	authorIds        []member.TenantMemberId
	ownerId          member.TenantMemberId
	createdAt        time.Time
	aggregateVersion uint
	eventMeta        aggregate.EventMeta
}

func (e ArtifactProposed) Id() values.TenantArtifactId {
	return e.id
}

func (e ArtifactProposed) Status() values.Status {
	return e.status
}

func (e ArtifactProposed) Name() values.Name {
	return e.name
}

func (e ArtifactProposed) Description() values.Description {
	return e.description
}

func (e ArtifactProposed) Url() values.Url {
	return e.url
}

func (e ArtifactProposed) CompetenceIds() []competence.TenantCompetenceId {
	return e.competenceIds[:]
}

func (e ArtifactProposed) AuthorIds() []member.TenantMemberId {
	return e.authorIds[:]
}

func (e ArtifactProposed) OwnerId() member.TenantMemberId {
	return e.ownerId
}

func (e ArtifactProposed) CreatedAt() time.Time {
	return e.createdAt
}

// EventType should be used instead of Invoke(Aggregate) approach
func (e ArtifactProposed) EventType() string {
	return "ArtifactProposed"
}

func (e ArtifactProposed) EventVersion() uint8 {
	return 1
}

func (e ArtifactProposed) AggregateVersion() uint {
	return e.aggregateVersion
}

func (e *ArtifactProposed) SetAggregateVersion(val uint) {
	e.aggregateVersion = val
}

func (e *ArtifactProposed) EventMeta() aggregate.EventMeta {
	return e.eventMeta
}

func (e *ArtifactProposed) SetEventMeta(val aggregate.EventMeta) {
	e.eventMeta = val
}

func (e ArtifactProposed) Export(ex ArtifactProposedExporterSetter) {
	ex.SetId(e.id)
	ex.SetStatus(e.status)
	ex.SetName(e.name)
	ex.SetDescription(e.description)
	ex.SetUrl(e.url)
	for i := range e.competenceIds {
		ex.AddCompetenceId(e.competenceIds[i])
	}
	for i := range e.authorIds {
		ex.AddAuthorId(e.authorIds[i])
	}
	ex.SetDescription(e.description)
	ex.SetOwnerId(e.ownerId)
	ex.SetCreatedAt(e.createdAt)
	ex.SetEventType(e.EventType())
	ex.SetEventVersion(e.EventVersion())
	ex.SetEventMeta(e.EventMeta())
	ex.SetAggregateVersion(e.AggregateVersion())
}

type ArtifactProposedExporterSetter interface {
	aggregate.PersistentDomainEventExporterSetter
	SetId(id values.TenantArtifactId)
	SetStatus(values.Status)
	SetName(values.Name)
	SetDescription(values.Description)
	SetUrl(values.Url)
	AddCompetenceId(competence.TenantCompetenceId)
	AddAuthorId(member.TenantMemberId)
	SetOwnerId(member.TenantMemberId)
	SetCreatedAt(time.Time)
}
