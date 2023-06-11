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

type DbSession interface {
	DbSessionExecutor
	DbSessionQuerier
}

type MutableQueryEvaluator interface {
	Evaluate(s DbSession) (Result, error)
}

// Deferred

type DeferredResultCallback func(Result)

type DeferredResult interface {
	AddCallback(DeferredResultCallback)
}

type DeferredDbSessionExecutor interface {
	Exec(query string, args ...any) (DeferredResult, error)
}

type DeferredDbSession interface {
	DeferredDbSessionExecutor
	DbSessionQuerier
}
