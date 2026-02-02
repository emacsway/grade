package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/aggregate"
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
	eventive  aggregate.EventiveEntity[aggregate.DomainEvent]
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
	ex.SetCreatedAt(t.createdAt)
	ex.SetVersion(t.Version())
}

type TenantExporterSetter interface {
	SetId(id values.TenantId)
	SetName(values.Name)
	SetCreatedAt(time.Time)
	SetVersion(uint)
}
