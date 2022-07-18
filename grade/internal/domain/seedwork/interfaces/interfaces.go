package interfaces

type Identity[T comparable] interface {
	Equals(Identity[T]) bool
	GetValue() T
}

type DomainEvent interface {
}

type EventiveEntity interface {
	AddDomainEvent(...DomainEvent)
	GetPendingDomainEvents() []DomainEvent
	ClearPendingDomainEvents()
}
