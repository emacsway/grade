package queries

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/grade/grade/internal/domain/specialist/endorsement"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

type SpecialistInsertQuery struct {
	params [9]any
}

func (q SpecialistInsertQuery) sql() string {
	return `
		INSERT INTO specialist (
			tenant_id, member_id, grade, created_at, version
		) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT DO NOTHING`
}

func (q *SpecialistInsertQuery) SetId(val memberVal.MemberId) {
	val.Export(q)
}

func (q *SpecialistInsertQuery) SetTenantId(val tenantVal.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *SpecialistInsertQuery) SetMemberId(val memberVal.InternalMemberId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *SpecialistInsertQuery) SetGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[2] = v
}

func (q *SpecialistInsertQuery) AddEndorsement(endorsement.Endorsement) {

}

func (q *SpecialistInsertQuery) AddAssignment(assignment.Assignment) {

}

func (q *SpecialistInsertQuery) SetCreatedAt(val time.Time) {
	q.params[3] = val
}

func (q *SpecialistInsertQuery) SetVersion(val uint) {
	q.params[4] = val
}

func (q *SpecialistInsertQuery) Evaluate(s session.DbSessionExecutor) (session.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
