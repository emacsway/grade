package seedwork

import "github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"

// There may be an event receiver here in the case of EventSourcing.

func NewEventiveEntity() (EventiveEntity, error) {
	return EventiveEntity{}, nil
}

type EventiveEntity struct {
	pendingDomainEvents []interfaces.DomainEvent
}

func (e EventiveEntity) AddDomainEvent(domainEvents ...interfaces.DomainEvent) {
	e.pendingDomainEvents = append(e.pendingDomainEvents, domainEvents...)
}

func (e EventiveEntity) GetPendingDomainEvents() []interfaces.DomainEvent {
	return e.pendingDomainEvents
}

func (e EventiveEntity) ClearPendingDomainEvents() {
	e.pendingDomainEvents = []interfaces.DomainEvent{}
}
