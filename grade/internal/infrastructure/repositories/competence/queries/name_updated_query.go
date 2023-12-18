package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

type NameUpdatedQuery struct {
	params [3]any
}

func (q NameUpdatedQuery) sql() string {
	return `
		UPDATE competence
		SET
			name = $3
		WHERE
			tenant_id = $1 AND competence_id = $2`
}

func (q *NameUpdatedQuery) SetId(val values.CompetenceId) {
	val.Export(q)
}

func (q *NameUpdatedQuery) SetTenantId(val tenantVal.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *NameUpdatedQuery) SetCompetenceId(val values.InternalCompetenceId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *NameUpdatedQuery) SetName(val values.Name) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[2] = v
}

func (q *NameUpdatedQuery) SetAggregateVersion(val uint) {
}

func (q *NameUpdatedQuery) SetEventType(val string) {
}

func (q *NameUpdatedQuery) Evaluate(s session.DbSession) (session.Result, error) {
	result, err := s.Exec(q.sql(), q.params[:]...)
	if err != nil {
		return result, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return result, err
	}
	if rowsAffected == 0 {
		return result, aggregate.ErrConcurrency
	}
	return result, err
}
