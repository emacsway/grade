package interfaces

import "github.com/emacsway/qualifying-grade/grade/pkg/domain/seedwork/interfaces"

type Exportable[T any] interface {
	Export() T
}

type Identity[T comparable] interface {
	Exportable[T]
	Equals(Identity[T]) bool
}

type PrimitiveExporter[T comparable] interface {
	SetState(T)
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
