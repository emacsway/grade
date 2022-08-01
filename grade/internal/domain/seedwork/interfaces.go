package seedwork

type DomainEvent interface {
}

type Session interface {
	Begin() (Session, error)
	Commit() error
	Rollback() error
}
