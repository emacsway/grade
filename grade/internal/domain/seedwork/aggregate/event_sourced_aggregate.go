package aggregate

import (
	"reflect"
	"strings"
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
		a.SetVersion(pastEvents[i].AggregateVersion())
		a.handlers[pastEvents[i].EventType()](pastEvents[i])
	}
}

func (a *EventSourcedAggregate) Update(e PersistentDomainEvent) {
	e.SetAggregateVersion(a.NextVersion())
	a.handlers[e.EventType()](e)
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

type PersistentDomainEventExporterSetter interface {
	SetEventType(string)
	SetEventVersion(uint8)
	SetEventMeta(EventMeta)
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
