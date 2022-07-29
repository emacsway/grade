package seedwork

type CompositeIdentifier[C any, D any] interface {
	Equals(CompositeIdentifier[C, D]) bool
	Export() C
	ExportTo(D)
}

type DomainEvent interface {
}

type Session interface {
	Begin() (Session, error)
	Commit() error
	Rollback() error
}
