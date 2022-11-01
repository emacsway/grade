package Endorsement

import (
	"fmt"
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type Params [9]any

type EndorsementInsertQuery struct {
	params []Params
}

func (q EndorsementInsertQuery) sql() string {
	sql := `
		INSERT INTO endorsement (
			tenant_id, specialist_id, specialist_grade, specialist_version,
			artifact_id, recognizer_id, recognizer_grade, recognizer_version, created_at
		) VALUES
		%s
		ON CONFLICT DO NOTHING`
	placeholders := `($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	bulkPlaceholders := ""

	for i := 0; i < len(q.params); i++ {
		if i != 0 {
			bulkPlaceholders += ", "
		}
		bulkPlaceholders += placeholders
	}
	return fmt.Sprintf(sql, bulkPlaceholders)
}

func (q EndorsementInsertQuery) flatParams() []any {
	var result []any
	for i := range q.params {
		result = append(result, q.params[i][:]...)
	}
	return result
}

func (q *EndorsementInsertQuery) SetSpecialistId(val member.TenantMemberId) {
	var v member.TenantMemberIdExporter
	val.Export(&v)
	q.addParam(0, v.TenantId)
	q.addParam(1, v.MemberId)
}

func (q *EndorsementInsertQuery) SetSpecialistGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.addParam(2, v)
}

func (q *EndorsementInsertQuery) SetSpecialistVersion(val uint) {
	q.addParam(3, val)
}

func (q *EndorsementInsertQuery) SetArtifactId(val artifact.TenantArtifactId) {
	var v artifact.TenantArtifactIdExporter
	val.Export(&v)
	q.addParam(4, v.ArtifactId)
}

func (q *EndorsementInsertQuery) SetRecognizerId(val member.TenantMemberId) {
	var v member.TenantMemberIdExporter
	val.Export(&v)
	q.addParam(5, v.MemberId)
}

func (q *EndorsementInsertQuery) SetRecognizerGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.addParam(6, v)
}

func (q *EndorsementInsertQuery) SetRecognizerVersion(val uint) {
	q.addParam(7, val)
}

func (q *EndorsementInsertQuery) SetCreatedAt(val time.Time) {
	q.addParam(8, val)
}

func (q *EndorsementInsertQuery) addParam(idx uint8, param any) {
	q.params[len(q.params)-1][idx] = param
}

func (q *EndorsementInsertQuery) Next() {
	q.params = append(q.params, Params{})
}

func (q *EndorsementInsertQuery) Execute(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {

	return s.Exec(q.sql(), q.flatParams()...)
}
