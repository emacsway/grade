package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	gradeVal "github.com/emacsway/grade/grade/internal/domain/specialist/assignment/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

type AssignmentInsertQuery struct {
	params [6]any
}

func (q AssignmentInsertQuery) sql() string {
	return `
		INSERT INTO specialist_assignment (
			tenant_id, specialist_id, specialist_version, assigned_grade, reason, created_at
		) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT DO NOTHING`
}

func (q *AssignmentInsertQuery) SetSpecialistId(val memberVal.MemberId) {
	var v memberVal.MemberIdExporter
	val.Export(&v)
	q.params[0] = v.TenantId
	q.params[1] = v.MemberId
}

func (q *AssignmentInsertQuery) SetSpecialistVersion(val uint) {
	q.params[2] = val
}

func (q *AssignmentInsertQuery) SetAssignedGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[3] = v
}

func (q *AssignmentInsertQuery) SetReason(val gradeVal.Reason) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[4] = v
}

func (q *AssignmentInsertQuery) SetCreatedAt(val time.Time) {
	q.params[5] = val
}

func (q *AssignmentInsertQuery) Evaluate(s session.DbSessionExecutor) (session.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
