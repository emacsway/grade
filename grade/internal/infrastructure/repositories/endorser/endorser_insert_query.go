package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type EndorserInsertQuery struct {
	params [7]any
}

func (q EndorserInsertQuery) sql() string {
	return `
		INSERT INTO endorser
		(tenant_id, member_id, grade, available_endorsement_count, pending_endorsement_count, version, created_at)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)`
}

func (q *EndorserInsertQuery) SetId(val member.TenantMemberId) {
	val.Export(q)
}

func (q *EndorserInsertQuery) SetTenantId(val tenant.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *EndorserInsertQuery) SetMemberId(val member.MemberId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *EndorserInsertQuery) SetGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[2] = v
}

func (q *EndorserInsertQuery) SetAvailableEndorsementCount(val endorser.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[3] = v
}

func (q *EndorserInsertQuery) SetPendingEndorsementCount(val endorser.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[4] = v
}

func (q *EndorserInsertQuery) SetVersion(val uint) {
	q.params[5] = val
}

func (q *EndorserInsertQuery) SetCreatedAt(val time.Time) {
	q.params[6] = val
}

func (q *EndorserInsertQuery) Evaluate(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
