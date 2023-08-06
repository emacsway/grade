package aggregate

type PersistentDomainEventHandler func(event PersistentDomainEvent)

func NewEventSourcedAggregate(version uint) EventSourcedAggregate {
	return EventSourcedAggregate{
		handlers:           make(map[string]PersistentDomainEventHandler),
		EventiveEntity:     EventiveEntity[PersistentDomainEvent]{},
		VersionedAggregate: NewVersionedAggregate(version),
	}
}

type EventSourcedAggregate struct {
	handlers map[string]PersistentDomainEventHandler
	EventiveEntity[PersistentDomainEvent]
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
func (a EventSourcedAggregate) PendingDomainEvents() []PersistentDomainEvent {
	return a.EventiveEntity.PendingDomainEvents()
}
