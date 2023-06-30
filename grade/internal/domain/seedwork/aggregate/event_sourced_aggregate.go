package aggregate

import (
	"reflect"
	"strings"
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewEventSourcedAggregate(version uint) EventSourcedAggregate {
	return EventSourcedAggregate{
		handlers:           make(map[string]PersistentDomainEventHandler),
		EventiveEntity:     EventiveEntity{},
		VersionedAggregate: NewVersionedAggregate(version),
	}
}

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

type PersistentDomainEventHandler func(event PersistentDomainEvent)

// The source of this data is domain layer.

type PersistentDomainEvent interface {
	DomainEvent
	EventType() string
	EventVersion() uint8
	EventMeta() EventMeta
	SetEventMeta(EventMeta)
	AggregateVersion() uint
	SetAggregateVersion(uint)
}

// The source of this data is application layer.

func NewEventMeta(
	eventId uuid.Uuid,
	correlationId uuid.Uuid,
	causationId uuid.Uuid,
	causalDependencies []CausalDependency,
	// occurredAt is the time of taking a slice of the state of the streams,
	// i.e. the time of obtaining the vector clock.
	// Therefore, it is the same for all aggregate events at the time of saving.
	occurredAt time.Time,
) EventMeta {
	return EventMeta{
		eventId:            eventId,
		correlationID:      correlationId,
		causationID:        causationId,
		causalDependencies: causalDependencies,
		occurredAt:         occurredAt,
	}
}

type EventMeta struct {
	eventId            uuid.Uuid
	correlationID      uuid.Uuid
	causationID        uuid.Uuid
	causalDependencies []CausalDependency
	occurredAt         time.Time
}

func (m EventMeta) EventId() uuid.Uuid {
	return m.eventId
}

func (m EventMeta) CorrelationId() uuid.Uuid {
	return m.correlationID
}

func (m EventMeta) CausationId() uuid.Uuid {
	return m.causationID
}

func (m EventMeta) CausalDependencies() []CausalDependency {
	return m.causalDependencies
}

func (m EventMeta) OccurredAt() time.Time {
	return m.occurredAt
}

func (m EventMeta) Spawn(eventId uuid.Uuid) EventMeta {
	n := m
	n.eventId = eventId
	return n
}

func (m EventMeta) Export(ex EventMetaExporterSetter) {
	ex.SetEventId(m.eventId)
	ex.SetCorrelationId(m.correlationID)
	ex.SetCausationId(m.causationID)
	for i := range m.causalDependencies {
		ex.AddCausalDependency(m.causalDependencies[i])
	}
	ex.SetOccurredAt(m.occurredAt)
}

type EventMetaExporterSetter interface {
	SetEventId(uuid.Uuid)
	SetCorrelationId(uuid.Uuid)
	SetCausationId(uuid.Uuid)
	AddCausalDependency(CausalDependency)
	SetOccurredAt(time.Time)
}

func NewCausalDependency(
	aggregateId any,
	aggregateType string,
	aggregateVersion uint,
) CausalDependency {
	return CausalDependency{
		aggregateId:      aggregateId,
		aggregateType:    aggregateType,
		aggregateVersion: aggregateVersion,
	}
}

type CausalDependency struct {
	aggregateId      any
	aggregateType    string
	aggregateVersion uint
}

func (d CausalDependency) AggregateId() any {
	return d.aggregateId
}

func (d CausalDependency) AggregateType() string {
	return d.aggregateType
}

func (d CausalDependency) AggregateVersion() uint {
	return d.aggregateVersion
}

func (d CausalDependency) Export(ex CausalDependencyExporterSetter) {
	ex.SetAggregateId(d.aggregateId)
	ex.SetAggregateType(d.aggregateType)
	ex.SetAggregateVersion(d.aggregateVersion)
}

type CausalDependencyExporterSetter interface {
	SetAggregateId(any)
	SetAggregateType(string)
	SetAggregateVersion(uint)
}

func BuildEventName(event DomainEvent) string {
	eventType := reflect.TypeOf(event).String()
	eventTypeParts := strings.Split(eventType, ".")
	eventName := eventTypeParts[len(eventTypeParts)-1]
	return eventName
}

func GetValueType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type() // .String()?
}
