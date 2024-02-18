package session

import "github.com/emacsway/grade/grade/internal/application/seedwork/session"

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Close() error
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Row interface {
	Err() error
	Scan(dest ...any) error
}

// Db

type DbSessionExecutor interface {
	Exec(query string, args ...any) (Result, error)
}

type DbSessionQuerier interface {
	Query(query string, args ...any) (Rows, error)
}

type DbSessionSingleQuerier interface {
	QueryRow(query string, args ...any) Row
}

type DbSession interface {
	session.Session
	DbSessionExecutor
	DbSessionQuerier
	DbSessionSingleQuerier
}

type QueryEvaluator interface {
	Evaluate(s DbSession) (Result, error)
}

type EventSourcedQueryEvaluator interface {
	QueryEvaluator
	SetStreamType(string)
}

// Deferred

type DeferredCallback[T interface{}] func(T) error

type Deferred[T interface{}] interface {
	Then(DeferredCallback[T]) error
	Catch(DeferredCallback[error]) error
}

type DeferredResult interface {
	Deferred[Result]
}

type DeferredRows interface {
	Then(DeferredCallback[Rows]) error
}

type DeferredRow interface {
	Then(DeferredCallback[Row]) error
}

type DeferredDbSessionExecutor interface {
	Exec(query string, args ...any) (DeferredResult, error)
}

type DeferredDbSessionQuerier interface {
	Query(query string, args ...any) (DeferredRows, error)
}

type DeferredDbSessionSingleQuerier interface {
	QueryRow(query string, args ...any) DeferredRow
}

type DeferredDbSession interface {
	DeferredDbSessionExecutor
	DeferredDbSessionQuerier
	DeferredDbSessionSingleQuerier
}
