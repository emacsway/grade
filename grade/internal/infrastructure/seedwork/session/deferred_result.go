package session

import (
	"errors"
)

type DeferredResultImp struct {
	lastInsertId int64
	rowsAffected int64
	callbacks    []DeferredResultCallback
	isResolved   bool
}

func (r *DeferredResultImp) Resolve(lastInsertId, rowsAffected int64) error {
	r.lastInsertId = lastInsertId
	r.rowsAffected = rowsAffected
	r.isResolved = true
	for i := range r.callbacks {
		err := r.callbacks[i](r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *DeferredResultImp) AddCallback(callback DeferredResultCallback) error {
	if r.isResolved {
		return callback(r)
	} else {
		r.callbacks = append(r.callbacks, callback)
	}
	return nil
}

func (r DeferredResultImp) LastInsertId() (int64, error) {
	if r.rowsAffected == 0 {
		return r.lastInsertId, nil
	} else {
		return 0, errors.New("LastInsertId is not supported by this driver")
	}
}

func (r DeferredResultImp) RowsAffected() (int64, error) {
	if r.lastInsertId == 0 {
		return r.rowsAffected, nil
	} else {
		return 0, errors.New("RowsAffected is not supported by INSERT command")
	}
}
