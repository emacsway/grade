package aggregate

type DomainEventAdder[T DomainEvent] interface {
	AddDomainEvent(...T)
}

type DomainEventAccessor[T DomainEvent] interface {
	PendingDomainEvents() []T
	ClearPendingDomainEvents()
}

func NewEventiveEntity() EventiveEntity[DomainEvent] {
	return EventiveEntity[DomainEvent]{}
}

type EventiveEntity[T DomainEvent] struct {
	pendingDomainEvents []T
}

func (e *EventiveEntity[T]) AddDomainEvent(domainEvents ...T) {
	e.pendingDomainEvents = append(e.pendingDomainEvents, domainEvents...)
}

func (e EventiveEntity[T]) PendingDomainEvents() []T {
	return e.pendingDomainEvents
}

func (e *EventiveEntity[T]) ClearPendingDomainEvents() {
	e.pendingDomainEvents = []T{}
}
