package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

func NewCompetence(
	id values.TenantCompetenceId,
	name values.Name,
	ownerId member.TenantMemberId,
	createdAt time.Time,
) (*Competence, error) {
	return &Competence{
		id:        id,
		name:      name,
		ownerId:   ownerId,
		createdAt: createdAt,
	}, nil
}

type Competence struct {
	id        values.TenantCompetenceId
	name      values.Name
	ownerId   member.TenantMemberId
	createdAt time.Time
	eventive  aggregate.EventiveEntity
	aggregate.VersionedAggregate
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
	ex.SetVersion(c.Version())
	ex.SetCreatedAt(c.createdAt)
}

type CompetenceExporterSetter interface {
	SetId(id values.TenantCompetenceId)
	SetName(values.Name)
	SetOwnerId(member.TenantMemberId)
	SetVersion(uint)
	SetCreatedAt(time.Time)
}
