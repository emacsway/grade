package seedwork

type Session interface {
	Atomic(func(Session) error) error
	/* Begin() (Session, error)
	Commit() error
	Rollback() error */
}
