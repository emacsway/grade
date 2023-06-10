package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type TenantInsertQuery struct {
	params   [3]any
	pkSetter func(any) error
}

func (q TenantInsertQuery) sql() string {
	return `
		INSERT INTO tenant
		(name, version, created_at)
		VALUES
		($1, $2, $3)
		RETURNING id`
}

func (q *TenantInsertQuery) SetId(val tenant.TenantId) {
	q.pkSetter = val.Scan
}

func (q *TenantInsertQuery) SetName(val tenant.Name) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *TenantInsertQuery) SetVersion(val uint) {
	q.params[1] = val
}

func (q *TenantInsertQuery) SetCreatedAt(val time.Time) {
	q.params[2] = val
}

func (q *TenantInsertQuery) Evaluate(s infrastructure.DbSessionInserter) (infrastructure.Result, error) {
	result, err := s.Insert(q.sql(), q.params[:]...)
	if err != nil {
		return result, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return result, err
	}
	err = q.pkSetter(id)
	return result, err
}
