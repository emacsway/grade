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

type RecognizerInsertQuery struct {
	params [7]any
}

func (q RecognizerInsertQuery) sql() string {
	return `
		INSERT INTO recognizer
		(tenant_id, member_id, grade, available_endorsement_count, pending_endorsement_count, version, created_at)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)`
}

func (q *RecognizerInsertQuery) SetId(val member.TenantMemberId) {
	val.Export(q)
}

func (q *RecognizerInsertQuery) SetTenantId(val tenant.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *RecognizerInsertQuery) SetMemberId(val member.MemberId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *RecognizerInsertQuery) SetGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[2] = v
}

func (q *RecognizerInsertQuery) SetAvailableEndorsementCount(val recognizer.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[3] = v
}

func (q *RecognizerInsertQuery) SetPendingEndorsementCount(val recognizer.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[4] = v
}

func (q *RecognizerInsertQuery) SetVersion(val uint) {
	q.params[5] = val
}

func (q *RecognizerInsertQuery) SetCreatedAt(val time.Time) {
	q.params[6] = val
}

func (q *RecognizerInsertQuery) Execute(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
