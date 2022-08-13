package seedwork

type Session interface {
	Atomic(func(Session) error) error
}
