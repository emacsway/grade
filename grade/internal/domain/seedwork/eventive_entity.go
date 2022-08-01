package seedwork

// There may be an event receiver here in the case of EventSourcing.

type EntityEventable interface {
	AddDomainEvent(...DomainEvent)
	PendingDomainEvents() []DomainEvent
	ClearPendingDomainEvents()
}

func NewEventiveEntity() EventiveEntity {
	return EventiveEntity{}
}

type EventiveEntity struct {
	pendingDomainEvents []DomainEvent
}

func (e EventiveEntity) AddDomainEvent(domainEvents ...DomainEvent) {
	e.pendingDomainEvents = append(e.pendingDomainEvents, domainEvents...)
}

func (e EventiveEntity) PendingDomainEvents() []DomainEvent {
	return e.pendingDomainEvents
}

func (e EventiveEntity) ClearPendingDomainEvents() {
	e.pendingDomainEvents = []DomainEvent{}
}
