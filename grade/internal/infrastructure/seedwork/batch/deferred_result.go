package batch

import (
	"errors"

	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

type DeferredResult struct {
	lastInsertId int64
	rowsAffected int64
	callbacks    []session.DeferredResultCallback
	isResolved   bool
}

func (r *DeferredResult) Resolve(lastInsertId, rowsAffected int64) {
	r.lastInsertId = lastInsertId
	r.rowsAffected = rowsAffected
	r.isResolved = true
	for i := range r.callbacks {
		r.callbacks[i](r)
	}
}

func (r *DeferredResult) SetRowsAffected(v int64) {
	r.rowsAffected = v
}

func (r *DeferredResult) AddCallback(callback session.DeferredResultCallback) {
	if r.isResolved {
		callback(r)
	} else {
		r.callbacks = append(r.callbacks, callback)
	}
}

func (r DeferredResult) LastInsertId() (int64, error) {
	if r.rowsAffected == 0 {
		return r.lastInsertId, nil
	} else {
		return 0, errors.New("LastInsertId is not supported by this driver")
	}
}

func (r DeferredResult) RowsAffected() (int64, error) {
	if r.lastInsertId == 0 {
		return r.rowsAffected, nil
	} else {
		return 0, errors.New("RowsAffected is not supported by INSERT command")
	}
}
