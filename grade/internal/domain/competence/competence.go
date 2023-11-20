package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/events"
	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

func NewCompetence(
	id values.CompetenceId,
	name values.Name,
	ownerId member.MemberId,
	createdAt time.Time,
) (*Competence, error) {
	agg := &Competence{
		id:        id,
		name:      name,
		ownerId:   ownerId,
		createdAt: createdAt,
	}
	e := events.NewCompetenceCreated(
		id,
		name,
		ownerId,
		createdAt,
	)
	e.SetAggregateVersion(agg.NextVersion())
	agg.eventive.AddDomainEvent(e)
	return agg, nil
}

type Competence struct {
	id        values.CompetenceId
	name      values.Name
	ownerId   member.MemberId
	createdAt time.Time
	eventive  aggregate.EventiveEntity[aggregate.DomainEvent]
	aggregate.VersionedAggregate
}

func (c *Competence) SetName(name values.Name) error {
	c.name = name
	e := events.NewNameUpdated(
		c.id,
		c.name,
	)
	e.SetAggregateVersion(c.NextVersion())
	c.eventive.AddDomainEvent(e)
	return nil
}

func (c Competence) PendingDomainEvents() []aggregate.DomainEvent {
	return c.eventive.PendingDomainEvents()
}

func (c *Competence) ClearPendingDomainEvents() {
	c.eventive.ClearPendingDomainEvents()
}

func (c Competence) Export(ex CompetenceExporterSetter) {
	ex.SetId(c.id)
	ex.SetName(c.name)
	ex.SetOwnerId(c.ownerId)
	ex.SetCreatedAt(c.createdAt)
	ex.SetVersion(c.Version())
}

type CompetenceExporterSetter interface {
	SetId(id values.CompetenceId)
	SetName(values.Name)
	SetOwnerId(member.MemberId)
	SetCreatedAt(time.Time)
	SetVersion(uint)
}
