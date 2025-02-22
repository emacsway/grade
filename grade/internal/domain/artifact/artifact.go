package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/events"
	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

func NewArtifact(
	id values.ArtifactId,
	status values.Status,
	name values.Name,
	description values.Description,
	url values.Url,
	competenceIds []competence.CompetenceId,
	authorIds []member.MemberId,
	ownerId member.MemberId,
	createdAt time.Time,
) (*Artifact, error) {
	e := events.NewArtifactProposed(
		id, status, name, description, url, competenceIds,
		authorIds, ownerId, createdAt,
	)
	agg, err := EmptyAggregate()
	if err != nil {
		return nil, err
	}
	agg.eventSourced.Update(e)
	return agg, nil
}

func EmptyAggregate() (*Artifact, error) {
	agg := &Artifact{
		eventSourced: aggregate.NewEventSourcedAggregate[aggregate.PersistentDomainEvent](0),
	}
	agg.eventSourced.AddHandler(&events.ArtifactProposed{}, agg.onArtifactProposed)
	return agg, nil
}

// Artifact is a good candidate for EventSourcing
type Artifact struct {
	id            values.ArtifactId
	status        values.Status
	name          values.Name
	description   values.Description
	url           values.Url
	competenceIds []competence.CompetenceId
	authorIds     []member.MemberId
	ownerId       member.MemberId
	createdAt     time.Time
	eventSourced  aggregate.EventSourcedAggregate[aggregate.PersistentDomainEvent]
}

func (a Artifact) Id() values.ArtifactId {
	return a.id
}

// TODO: Use Specification pattern instead?
// https://enterprisecraftsmanship.com/posts/specification-pattern-always-valid-domain-model/

func (a Artifact) HasAuthor(authorId member.MemberId) bool {
	for i := range a.authorIds {
		if a.authorIds[i].Equal(authorId) {
			return true
		}
	}
	return false
}

func (a Artifact) PendingDomainEvents() []aggregate.PersistentDomainEvent {
	return a.eventSourced.PendingDomainEvents()
}

func (a *Artifact) ClearPendingDomainEvents() {
	a.eventSourced.ClearPendingDomainEvents()
}

func (a Artifact) Version() uint {
	return a.eventSourced.Version()
}

func (a *Artifact) SetVersion(val uint) {
	a.eventSourced.SetVersion(val)
}

func (a Artifact) Export(ex ArtifactExporterSetter) {
	ex.SetId(a.id)
	ex.SetStatus(a.status)
	ex.SetName(a.name)
	ex.SetDescription(a.description)
	ex.SetUrl(a.url)
	for i := range a.competenceIds {
		ex.AddCompetenceId(a.competenceIds[i])
	}
	for i := range a.authorIds {
		ex.AddAuthorId(a.authorIds[i])
	}
	ex.SetDescription(a.description)
	ex.SetOwnerId(a.ownerId)
	ex.SetCreatedAt(a.createdAt)
	ex.SetVersion(a.Version())
}

func (a *Artifact) onArtifactProposed(e aggregate.PersistentDomainEvent) {
	et := e.(*events.ArtifactProposed)
	a.id = et.AggregateId()
	a.status = et.Status()
	a.name = et.Name()
	a.description = et.Description()
	a.url = et.Url()
	a.competenceIds = et.CompetenceIds()
	a.authorIds = et.AuthorIds()
	a.ownerId = et.OwnerId()
	a.createdAt = et.CreatedAt()
}

type ArtifactExporterSetter interface {
	SetId(id values.ArtifactId)
	SetStatus(values.Status)
	SetName(values.Name)
	SetDescription(values.Description)
	SetUrl(values.Url)
	AddCompetenceId(competence.CompetenceId)
	AddAuthorId(member.MemberId)
	SetOwnerId(member.MemberId)
	SetCreatedAt(time.Time)
	SetVersion(uint)
}
