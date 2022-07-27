package interfaces

type Exporter[T any] interface {
	SetState(T)
}

type ExportableTo[T any] interface {
	ExportTo(Exporter[T])
}

// alternative approach:

type Exportable[T any] interface {
	Export() T
}

type Identity[T comparable] interface {
	Exportable[T]
	ExportableTo[T]
	Equals(Identity[T]) bool
}

type CompositeIdentity[C any, D any] interface {
	Equals(CompositeIdentity[C, D]) bool
	Export() C
	ExportTo(D)
}

type DomainEvent interface {
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
