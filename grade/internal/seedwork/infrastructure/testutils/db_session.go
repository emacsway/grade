package testutils

import (
	"database/sql"
	"errors"

	appSession "github.com/emacsway/grade/grade/internal/seedwork/application/session"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

func NewDbSessionStub(rows *RowsStub) *DbSessionStub {
	return &DbSessionStub{
		Rows: rows,
	}
}

type DbSessionStub struct {
	Rows         *RowsStub
	ActualQuery  string
	ActualParams []any
}

func (s DbSessionStub) Atomic(callback appSession.SessionCallback) error {
	return callback(s)
}

func (s *DbSessionStub) Exec(query string, args ...any) (session.Result, error) {
	s.ActualQuery = query
	s.ActualParams = args
	return session.NewDeferredResult(), nil
}

func (s *DbSessionStub) Query(query string, args ...any) (session.Rows, error) {
	s.ActualQuery = query
	s.ActualParams = args
	return s.Rows, nil
}

func (s *DbSessionStub) QueryRow(query string, args ...any) session.Row {
	s.ActualQuery = query
	s.ActualParams = args
	return s.Rows
}

func NewRowsStub(rows ...[]any) *RowsStub {
	return &RowsStub{
		rows, 0, false,
	}
}

type RowsStub struct {
	rows   [][]any
	idx    int
	Closed bool
}

func (r *RowsStub) Close() error {
	r.Closed = true
	return nil
}

func (r RowsStub) Err() error {
	return nil
}

func (r *RowsStub) Next() bool {
	r.idx++
	return len(r.rows) < r.idx
}

func (r RowsStub) Scan(dest ...any) error {
	for i, d := range dest {
		dt, ok := d.(sql.Scanner)
		if !ok {
			return errors.New("value should implement sql.Scanner interface")
		}
		err := dt.Scan(r.rows[r.idx][i])
		if err != nil {
			return err
		}
	}
	return nil
}

type RowStub struct {
	rows *RowsStub
}

func (r *RowStub) Err() error {
	return r.rows.Err()
}

func (r *RowStub) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}
