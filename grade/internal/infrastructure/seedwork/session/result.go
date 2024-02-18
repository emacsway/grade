package session

import (
	"errors"
)

func NewResult(lastInsertId, rowsAffected int64) *DeferredResultImp {
	r := &DeferredResultImp{}
	err := r.Resolve(lastInsertId, rowsAffected)
	if err != nil {
		panic(err)
	}
	return r
}

func NewDeferredResult() *DeferredResultImp {
	return &DeferredResultImp{}
}

type ResultImp struct {
	lastInsertId int64
	rowsAffected int64
}

func (r ResultImp) LastInsertId() (int64, error) {
	if r.rowsAffected == 0 {
		return r.lastInsertId, nil
	} else {
		return 0, errors.New("LastInsertId is not supported by this driver")
	}
}

func (r ResultImp) RowsAffected() (int64, error) {
	if r.lastInsertId == 0 {
		return r.rowsAffected, nil
	} else {
		return 0, errors.New("RowsAffected is not supported by INSERT command")
	}
}

type DeferredResultImp struct {
	DeferredImp[Result]
	ResultImp
}

func (r *DeferredResultImp) Resolve(lastInsertId, rowsAffected int64) error {
	r.ResultImp = ResultImp{lastInsertId, rowsAffected}
	return r.DeferredImp.Resolve(r)
}
