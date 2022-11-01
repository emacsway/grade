package recognizer

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/recognizer"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type RecognizerUpdateQuery struct {
	params [7]any
}

func (q RecognizerUpdateQuery) sql() string {
	return `
		UPDATE recognizer SET
			grade = $4,
			available_endorsement_count = $5,
			pending_endorsement_count = $6,
			version = version + 1
		WHERE
			tenant_id = $1, member_id=$2, version = $3`
}

func (q *RecognizerUpdateQuery) SetId(val member.TenantMemberId) {
	val.Export(q)
}

func (q *RecognizerUpdateQuery) SetTenantId(val tenant.TenantId) {
	var v exporters.UuidExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *RecognizerUpdateQuery) SetMemberId(val member.MemberId) {
	var v exporters.UuidExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *RecognizerUpdateQuery) SetGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[2] = v
}

func (q *RecognizerUpdateQuery) SetAvailableEndorsementCount(val recognizer.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[3] = v
}

func (q *RecognizerUpdateQuery) SetPendingEndorsementCount(val recognizer.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[4] = v
}

func (q *RecognizerUpdateQuery) SetVersion(val uint) {
	q.params[5] = val
}

func (q *RecognizerUpdateQuery) SetCreatedAt(val time.Time) {
	q.params[6] = val
}

func (q *RecognizerUpdateQuery) Execute(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
