package seedwork

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TenantId interface {
	Exportable[uint64]
}

type CausalDependency interface {
	TenantId() uint64
	StreamType() string
	StreamID() string
	StreamPosition() uint
}

type EventMeta interface {
	EventId() uuid.UUID
	CorrelationID() uuid.UUID
	CausationID() uuid.UUID
	CausalDependencies() []CausalDependency
	OccurredAt() time.Time
}

type PersistentDomainEvent interface {
	DomainEvent
	EventType() string
	EventVersion() uint8
	TenantId() TenantId // To be able to drop quickly the whole tenant
	SetTenantId(TenantId)
	StreamType() string // For Causal Dependencies reason and for variability of retention policy
	SetStreamType(string)
	StreamId() fmt.Stringer
	SetStreamId(value fmt.Stringer)
	StreamPosition() uint
	SetStreamPosition(uint)
}

type PersistentDomainEventHandler func(event PersistentDomainEvent)

func NewEventSourcedAggregate(
	tenantId TenantId,
	streamId fmt.Stringer,
	streamType string,
) EventSourcedAggregate {
	return EventSourcedAggregate{
		tenantId:   tenantId,
		streamId:   streamId,
		streamType: streamType,
	}
}

type EventSourcedAggregate struct {
	tenantId   TenantId
	streamId   fmt.Stringer
	streamType string
	handlers   map[string]PersistentDomainEventHandler
	EventiveEntity
	VersionedAggregate
}

func (a *EventSourcedAggregate) AddHandler(e PersistentDomainEvent, handler PersistentDomainEventHandler) {
	a.handlers[e.EventType()] = handler
}

func (a *EventSourcedAggregate) LoadFrom(pastEvents []PersistentDomainEvent) {
	for i := range pastEvents {
		a.handlers[pastEvents[i].EventType()](pastEvents[i])
		a.SetVersion(pastEvents[i].StreamPosition())
	}
}

func (a *EventSourcedAggregate) Update(e PersistentDomainEvent) {
	e.SetTenantId(a.tenantId)
	e.SetStreamId(a.streamId)
	e.SetStreamType(a.streamType)
	e.SetStreamPosition(a.Version() + 1)
	a.handlers[e.EventType()](e)
	a.SetVersion(a.Version() + 1)
	a.AddDomainEvent(e)
}
