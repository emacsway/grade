package infrastructure

import (
	"strings"

	"database/sql"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/emacsway/grade/grade/internal/application"
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

func (s *PgxSession) Atomic(callback func(session application.Session) error) error {
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
	if IsInsertQuery(query) {
		var id int
		err := s.dbExecutor.QueryRow(query, args...).Scan(&id)
		return LastInsertId(id), err
	}
	return s.dbExecutor.Exec(query, args...)
}

func (s *PgxSession) Insert(query string, args ...any) (Result, error) {
	var id int
	err := s.dbExecutor.QueryRow(query, args...).Scan(&id)
	return LastInsertId(id), err
}

func (s *PgxSession) Query(query string, args ...any) (Rows, error) {
	return s.dbExecutor.Query(query, args...)
}

type DbExecutor interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type LastInsertId int64

func (i LastInsertId) LastInsertId() (int64, error) {
	return int64(i), nil
}

func (i LastInsertId) RowsAffected() (int64, error) {
	return 0, errors.New("RowsAffected is not supported by INSERT command")
}

type RowsAffected int64

func (RowsAffected) LastInsertId() (int64, error) {
	return 0, errors.New("LastInsertId is not supported by this driver")
}

func (v RowsAffected) RowsAffected() (int64, error) {
	return int64(v), nil
}

func IsInsertQuery(sql string) bool {
	return strings.TrimSpace(sql)[:6] == "INSERT"
}
