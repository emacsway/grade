package testutils

import (
	"database/sql"
	"errors"

	"github.com/emacsway/grade/grade/internal/application"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
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

func (s DbSessionStub) Atomic(callback application.SessionCallback) error {
	return callback(s)
}

func (s *DbSessionStub) Exec(query string, args ...any) (session.Result, error) {
	s.ActualQuery = query
	s.ActualParams = args
	return &ResultStub{}, nil
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

type ResultStub struct {
	lastInsertId int64
	rowsAffected int64
	callbacks    []session.DeferredResultCallback
}

func (r *ResultStub) Resolve(lastInsertId, rowsAffected int64) {
	r.lastInsertId = lastInsertId
	r.rowsAffected = rowsAffected
	for i := range r.callbacks {
		r.callbacks[i](r)
	}
}

func (r *ResultStub) SetRowsAffected(v int64) {
	r.rowsAffected = v
}

func (r *ResultStub) AddCallback(callback session.DeferredResultCallback) {
	r.callbacks = append(r.callbacks, callback)
}

func (r ResultStub) LastInsertId() (int64, error) {
	if r.rowsAffected == 0 {
		return r.lastInsertId, nil
	} else {
		return 0, errors.New("LastInsertId is not supported by this driver")
	}
}

func (r ResultStub) RowsAffected() (int64, error) {
	if r.lastInsertId == 0 {
		return r.rowsAffected, nil
	} else {
		return 0, errors.New("RowsAffected is not supported by INSERT command")
	}
}
