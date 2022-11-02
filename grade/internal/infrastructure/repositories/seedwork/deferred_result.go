package seedwork

import (
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type DeferredResult struct {
	lastInsertId int64
	rowsAffected int64
	callbacks    []infrastructure.DeferredResultCallback
}

func (r *DeferredResult) Resolve(lastInsertId int64, rowsAffected int64) {
	r.lastInsertId = lastInsertId
	r.rowsAffected = rowsAffected
	for i := range r.callbacks {
		r.callbacks[i](r)
	}
}

func (r *DeferredResult) SetRowsAffected(v int64) {
	r.rowsAffected = v
}

func (r *DeferredResult) AddCallback(callback infrastructure.DeferredResultCallback) {
	r.callbacks = append(r.callbacks, callback)
}

func (r DeferredResult) LastInsertId() (int64, error) {
	return r.lastInsertId, nil
}
func (r DeferredResult) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}
