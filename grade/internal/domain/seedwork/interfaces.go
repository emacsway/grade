package seedwork

type Session interface {
	Begin() (Session, error)
	Commit() error
	Rollback() error
}
