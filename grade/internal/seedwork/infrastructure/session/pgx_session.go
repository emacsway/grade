package session

import (
	"strings"

	"database/sql"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/emacsway/grade/grade/internal/seedwork/application/session"
)

func NewPgxSession(db *sql.DB) *PgxSession {
	return &PgxSession{
		db:         db,
		dbExecutor: db,
	}
}

type PgxSession struct {
	db         *sql.DB
	dbExecutor DbExecutor
}

func (s *PgxSession) Atomic(callback session.SessionCallback) error {
	// TODO: Add support for SavePoint:
	// https://github.com/golang/go/issues/7898#issuecomment-580080390
	if s.db == nil {
		return errors.New("savePoint is not currently supported")
	}
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}
	newSession := &PgxSession{
		dbExecutor: tx,
	}
	err = callback(newSession)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return multierror.Append(err, txErr)
		}
		return err
	}
	if txErr := tx.Commit(); txErr != nil {
		return errors.Wrap(err, "failed to commit tx")
	}
	return nil
}

func (s *PgxSession) Exec(query string, args ...any) (Result, error) {
	if IsAutoincrementInsertQuery(query) {
		return s.insert(query, args...)
	}
	return s.dbExecutor.Exec(query, args...)
}

func (s *PgxSession) insert(query string, args ...any) (Result, error) {
	var id int64
	err := s.dbExecutor.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return nil, err
	}
	return NewResult(id, 0), nil
}

func (s *PgxSession) Query(query string, args ...any) (Rows, error) {
	return s.dbExecutor.Query(query, args...)
}

func (s *PgxSession) QueryRow(query string, args ...any) Row {
	return s.dbExecutor.QueryRow(query, args...)
}

type DbExecutor interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

func IsInsertQuery(query string) bool {
	return strings.TrimSpace(query)[:6] == "INSERT" && !strings.Contains(query, "RETURNING")
}

func IsAutoincrementInsertQuery(query string) bool {
	return strings.TrimSpace(query)[:6] == "INSERT" && strings.Contains(query, "RETURNING")
}
