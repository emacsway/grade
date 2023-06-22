package infrastructure

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
	DbSessionExecutor
	DbSessionQuerier
	DbSessionSingleQuerier
}

type QueryEvaluator interface {
	Evaluate(s DbSession) (Result, error)
}

// Deferred

type DeferredResultCallback func(Result)

type DeferredResult interface {
	AddCallback(DeferredResultCallback)
}

type DeferredRowsCallback func(Rows)

type DeferredRows interface {
	AddCallback(DeferredRowsCallback)
}

type DeferredRowCallback func(Rows)

type DeferredRow interface {
	AddCallback(DeferredRowsCallback)
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
