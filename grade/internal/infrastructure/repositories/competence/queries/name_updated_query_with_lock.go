package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/aggregate"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

/**
 * See how to update multiple tables (entity without version column and root entity with it) at once:
 *
 * - https://dba.stackexchange.com/a/215655
 * - https://www.postgresql.org/docs/current/queries-with.html#QUERIES-WITH-MODIFYING
 *
 *     with shape_update as (
 *       UPDATE prd_shape
 *        SET name_en = 'Another name'
 *       WHERE serial_number = '1234ST'
 *       returning id, serial_number
 *     )
 *     UPDATE prd_sectionshapename
 *       SET company_shape_name = 'ABC'
 *     WHERE (shape_id, serial_number) IN (select id, serial_number from shape_update);
**/

type NameUpdatedQueryWithLock struct {
	params [4]any
}

func (q NameUpdatedQueryWithLock) sql() string {
	return `
		UPDATE competence
		SET
			name = $4,
			version = $3
		WHERE
			tenant_id = $1 AND competence_id = $2 AND version = $3-1`
}

func (q *NameUpdatedQueryWithLock) SetId(val values.CompetenceId) {
	val.Export(q)
}

func (q *NameUpdatedQueryWithLock) SetTenantId(val tenantVal.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *NameUpdatedQueryWithLock) SetCompetenceId(val values.InternalCompetenceId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *NameUpdatedQueryWithLock) SetAggregateVersion(val uint) {
	q.params[2] = val
}

func (q *NameUpdatedQueryWithLock) SetName(val values.Name) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[3] = v
}

func (q *NameUpdatedQueryWithLock) SetEventType(val string) {
}

func (q *NameUpdatedQueryWithLock) Evaluate(s session.DbSession) (session.Result, error) {
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
