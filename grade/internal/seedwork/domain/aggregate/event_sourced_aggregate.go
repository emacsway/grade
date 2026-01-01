package aggregate

type PersistentDomainEventHandler func(event PersistentDomainEvent)

func NewEventSourcedAggregate[T PersistentDomainEvent](version uint) EventSourcedAggregate[T] {
	return EventSourcedAggregate[T]{
		handlers:           make(map[string]PersistentDomainEventHandler),
		EventiveEntity:     EventiveEntity[T]{},
		VersionedAggregate: NewVersionedAggregate(version),
	}
}

type EventSourcedAggregate[T PersistentDomainEvent] struct {
	handlers map[string]PersistentDomainEventHandler
	EventiveEntity[T]
	VersionedAggregate
}

func (a *EventSourcedAggregate[T]) AddHandler(e T, handler PersistentDomainEventHandler) {
	a.handlers[e.EventType()] = handler
}

func (a *EventSourcedAggregate[T]) LoadFrom(pastEvents []T) {
	for i := range pastEvents {
		a.SetVersion(pastEvents[i].AggregateVersion())
		a.handlers[pastEvents[i].EventType()](pastEvents[i])
	}
}

func (a *EventSourcedAggregate[T]) Update(e T) {
	e.SetAggregateVersion(a.NextVersion())
	a.AddDomainEvent(e)
	a.handlers[e.EventType()](e)
}
func (a EventSourcedAggregate[T]) PendingDomainEvents() []T {
	return a.EventiveEntity.PendingDomainEvents()
}
