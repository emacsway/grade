package queries

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type CompetenceCreatedQuery struct {
	params   [5]any
	pkSetter func(any) error
}

func (q CompetenceCreatedQuery) sql() string {
	return `
		INSERT INTO competence
		(tenant_id, name, owner_id, created_at, version)
		VALUES
		($1, $2, $3, $4, $5)
		RETURNING competence_id`
}

func (q *CompetenceCreatedQuery) SetId(val values.TenantCompetenceId) {
	val.Export(q)
}

func (q *CompetenceCreatedQuery) SetTenantId(val tenantVal.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *CompetenceCreatedQuery) SetCompetenceId(val values.CompetenceId) {
	q.pkSetter = val.Scan
}

func (q *CompetenceCreatedQuery) SetName(val values.Name) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *CompetenceCreatedQuery) SetOwnerId(val memberVal.MemberId) {
	var v memberVal.MemberIdExporter
	val.Export(&v)
	q.params[2] = v.MemberId
}

func (q *CompetenceCreatedQuery) SetCreatedAt(val time.Time) {
	q.params[3] = val
}

func (q *CompetenceCreatedQuery) SetAggregateVersion(val uint) {
	q.params[4] = val
}

func (q *CompetenceCreatedQuery) SetEventType(val string) {
}

func (q *CompetenceCreatedQuery) Evaluate(s infrastructure.DbSession) (infrastructure.Result, error) {
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
