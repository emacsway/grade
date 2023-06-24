package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewTenant(
	id values.TenantId,
	name values.Name,
	createdAt time.Time,
) (*Tenant, error) {
	return &Tenant{
		id:        id,
		name:      name,
		createdAt: createdAt,
	}, nil
}

type Tenant struct {
	id        values.TenantId
	name      values.Name
	createdAt time.Time
	eventive  aggregate.EventiveEntity
	aggregate.VersionedAggregate
}

func (t Tenant) PendingDomainEvents() []aggregate.DomainEvent {
	return t.eventive.PendingDomainEvents()
}

func (t *Tenant) ClearPendingDomainEvents() {
	t.eventive.ClearPendingDomainEvents()
}

func (t Tenant) Export(ex TenantExporterSetter) {
	ex.SetId(t.id)
	ex.SetName(t.name)
	ex.SetVersion(t.Version())
	ex.SetCreatedAt(t.createdAt)
}

type TenantExporterSetter interface {
	SetId(id values.TenantId)
	SetName(values.Name)
	SetVersion(uint)
	SetCreatedAt(time.Time)
}
