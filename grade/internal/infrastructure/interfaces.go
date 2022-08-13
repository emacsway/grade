package infrastructure

import (
	"github.com/emacsway/grade/grade/internal/application"
)

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Close() error
	Err() error
	Next() bool
}

type DbSession interface {
	application.Session

	Exec(query string, args ...any) (Result, error)
	Fetch(query string, args ...any) (Rows, error)
}
