package interfaces

import "github.com/emacsway/qualifying-grade/grade/pkg/domain/seedwork/interfaces"

type Originator[T any] interface {
	CreateMemento() T
	// SetMemento(T)
}

type Identity[T comparable] interface {
	Originator[T]
	Equals(Identity[T]) bool
}

// Public listeners is a subset of all listeners, not vice versa. Suppose, we have three circles...

type DomainEvent interface {
	interfaces.PublicDomainEvent
}

type EventiveEntity interface {
	AddDomainEvent(...DomainEvent)
	GetPendingDomainEvents() []DomainEvent
	ClearPendingDomainEvents()
}

type VersionedAggregate interface {
	GetVersion() uint
	IncreaseVersion()
}

type Session interface {
	Begin() (Session, error)
	Commit() error
	Rollback() error
}
