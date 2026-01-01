package aggregate

import (
	"reflect"
	"strings"
)

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
