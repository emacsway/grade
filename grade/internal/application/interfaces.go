package application

type Session interface {
	Atomic(func(Session) error) error
}
