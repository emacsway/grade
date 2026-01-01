package queries

import (
	"time"

	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

type TenantInsertQuery struct {
	params   [3]any
	pkSetter func(any) error
}

func (q TenantInsertQuery) sql() string {
	return `
		INSERT INTO tenant
		(name, created_at, version)
		VALUES
		($1, $2, $3)
		RETURNING id`
}

func (q *TenantInsertQuery) SetId(val tenantVal.TenantId) {
	q.pkSetter = val.Scan
}

func (q *TenantInsertQuery) SetName(val tenantVal.Name) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *TenantInsertQuery) SetCreatedAt(val time.Time) {
	q.params[1] = val
}

func (q *TenantInsertQuery) SetVersion(val uint) {
	q.params[2] = val
}

func (q *TenantInsertQuery) Evaluate(s session.DbSessionExecutor) (session.Result, error) {
	result, err := s.Exec(q.sql(), q.params[:]...)
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
