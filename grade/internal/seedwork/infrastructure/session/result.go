package session

import (
	"errors"

	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/deferred"
)

func NewResult(lastInsertId, rowsAffected int64) *DeferredResultImp {
	r := &DeferredResultImp{}
	r.Resolve(lastInsertId, rowsAffected)
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
	ResultImp
	deferred.DeferredImp[Result]
}

func (r *DeferredResultImp) Resolve(lastInsertId, rowsAffected int64) {
	r.ResultImp = ResultImp{lastInsertId, rowsAffected}
	r.DeferredImp.Resolve(r)
}
