package seedwork

import (
	"time"

	"github.com/google/uuid"
)

type CausalDependency interface {
	AggregateId() interface{}
	AggregateType() string
	AggregateVersion() uint
}

// The source this data is application layer.

type EventMeta interface {
	EventId() uuid.UUID
	CorrelationID() uuid.UUID
	CausationID() uuid.UUID
	CausalDependencies() []CausalDependency
	OccurredAt() time.Time
}

// The source this data is domain layer.

type PersistentDomainEvent interface {
	DomainEvent
	EventType() string
	EventVersion() uint8
	AggregateVersion() uint
	SetAggregateVersion(uint)
}

type PersistentDomainEventHandler func(event PersistentDomainEvent)

type EventSourcedAggregate struct {
	handlers map[string]PersistentDomainEventHandler
	EventiveEntity
	VersionedAggregate
}

func (a *EventSourcedAggregate) AddHandler(e PersistentDomainEvent, handler PersistentDomainEventHandler) {
	a.handlers[e.EventType()] = handler
}

func (a *EventSourcedAggregate) LoadFrom(pastEvents []PersistentDomainEvent) {
	for i := range pastEvents {
		a.handlers[pastEvents[i].EventType()](pastEvents[i])
		a.SetVersion(pastEvents[i].AggregateVersion())
	}
}

func (a *EventSourcedAggregate) Update(e PersistentDomainEvent) {
	e.SetAggregateVersion(a.Version() + 1)
	a.handlers[e.EventType()](e)
	a.SetVersion(a.Version() + 1)
	a.AddDomainEvent(e)
}
