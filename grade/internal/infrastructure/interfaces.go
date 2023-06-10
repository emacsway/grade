package infrastructure

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type DeferredResultCallback func(Result)

type DeferredResult interface {
	AddCallback(DeferredResultCallback)
}

type Rows interface {
	Close() error
	Err() error
	Next() bool
}

type DbSessionExecutor interface {
	Exec(query string, args ...any) (Result, error)
}

type DbSessionInserter interface {
	Insert(query string, args ...any) (Result, error)
}

type DbSessionQuerier interface {
	Query(query string, args ...any) (Rows, error)
}

type DbSession interface {
	DbSessionExecutor
	DbSessionInserter
	DbSessionQuerier
}

type MutableQueryEvaluator interface {
	Evaluate(s DbSessionExecutor) (Result, error)
}
