package interfaces

type Exporter[T any] interface {
	SetState(T)
}

type Exporting[T any] interface {
	ExportTo(Exporter[T])
}

// alternative approach:

type Exportable[T any] interface {
	Export() T
}

type Identity[T comparable] interface {
	Exportable[T]
	Exporting[T]
	Equals(Identity[T]) bool
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
