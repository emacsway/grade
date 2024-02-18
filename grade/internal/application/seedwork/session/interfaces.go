package session

import "context"

type SessionCallback func(Session) error

type Session interface {
	Atomic(SessionCallback) error
}

type SessionContextCallback func(SessionContext) error

type SessionContext interface {
	context.Context
	Atomic(SessionContextCallback) error
}
