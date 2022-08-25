package assignment

import (
	"fmt"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"time"
)

type Params [6]any

type AssignmentInsertQuery struct {
	params []Params
}

func (q AssignmentInsertQuery) sql() string {
	sql := `
		INSERT INTO assignment (
			tenant_id, specialist_id, specialist_version, assigned_grade, created_at, reason
		) VALUES
		%s
		ON CONFLICT DO NOTHING`
	placeholders := `($1, $2, $3, $4, $5, $6)`

	bulkPlaceholders := ""

	for i := 0; i < len(q.params); i++ {
		if i != 0 {
			bulkPlaceholders += ", "
		}
		bulkPlaceholders += placeholders
	}
	return fmt.Sprintf(sql, bulkPlaceholders)
}

func (q AssignmentInsertQuery) flatParams() []any {
	var result []any
	for i := range q.params {
		result = append(result, q.params[i][:]...)
	}
	return result
}

func (q *AssignmentInsertQuery) SetSpecialistId(val member.TenantMemberId) {
	var v member.TenantMemberIdExporter
	val.Export(&v)
	q.addParam(0, v.TenantId)
	q.addParam(1, v.MemberId)
}

func (q *AssignmentInsertQuery) SetSpecialistVersion(val uint) {
	q.addParam(2, val)
}

func (q *AssignmentInsertQuery) SetAssignedGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.addParam(3, val)
}

func (q *AssignmentInsertQuery) SetReason(val assignment.Reason) {
	var v exporters.StringExporter
	val.Export(&v)
	q.addParam(4, val)
}

func (q *AssignmentInsertQuery) SetCreatedAt(val time.Time) {
	q.addParam(5, val)
}

func (q *AssignmentInsertQuery) addParam(idx uint8, param any) {
	q.params[len(q.params)-1][idx] = param
}

func (q *AssignmentInsertQuery) Next() {
	q.params = append(q.params, Params{})
}

func (q *AssignmentInsertQuery) Execute(s infrastructure.DbSession) (infrastructure.Result, error) {
	return s.Exec(q.sql(), q.flatParams()...)
}
