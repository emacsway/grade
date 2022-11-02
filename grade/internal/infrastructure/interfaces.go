package infrastructure

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Close() error
	Err() error
	Next() bool
}

type DbSessionExecutor interface {
	Exec(query string, args ...any) (Result, error)
}

type DbSessionFetcher interface {
	Fetch(query string, args ...any) (Rows, error)
}

type MutableQuery interface {
	Execute(s DbSessionExecutor) (Result, error)
}
