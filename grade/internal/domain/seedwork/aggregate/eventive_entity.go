package aggregate

type EventiveEntityAccessor interface {
	AddDomainEvent(...DomainEvent)
	PendingDomainEvents() []DomainEvent
	ClearPendingDomainEvents()
}

type EventiveEntityAdder interface {
	AddDomainEvent(...DomainEvent)
}

type EventiveEntity struct {
	pendingDomainEvents []DomainEvent
}

func (e *EventiveEntity) AddDomainEvent(domainEvents ...DomainEvent) {
	e.pendingDomainEvents = append(e.pendingDomainEvents, domainEvents...)
}

func (e EventiveEntity) PendingDomainEvents() []DomainEvent {
	return e.pendingDomainEvents
}

func (e *EventiveEntity) ClearPendingDomainEvents() {
	e.pendingDomainEvents = []DomainEvent{}
}
