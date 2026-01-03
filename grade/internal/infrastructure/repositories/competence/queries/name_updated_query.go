package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/aggregate"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
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
	val.Export(func(v uint) { q.params[0] = v })
}

func (q *NameUpdatedQuery) SetCompetenceId(val values.InternalCompetenceId) {
	val.Export(func(v uint) { q.params[1] = v })
}

func (q *NameUpdatedQuery) SetName(val values.Name) {
	val.Export(func(v string) { q.params[2] = v })
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
