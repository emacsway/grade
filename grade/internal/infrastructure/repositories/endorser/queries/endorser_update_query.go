package queries

import (
	"time"

	endorserVal "github.com/emacsway/grade/grade/internal/domain/endorser/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type EndorserUpdateQuery struct {
	params [7]any
}

func (q EndorserUpdateQuery) sql() string {
	return `
		UPDATE endorser SET
			grade = $4,
			available_endorsement_count = $5,
			pending_endorsement_count = $6,
			version = version + 1
		WHERE
			tenant_id = $1 AND member_id = $2 AND version = $3`
}

func (q *EndorserUpdateQuery) SetId(val member.TenantMemberId) {
	val.Export(q)
}

func (q *EndorserUpdateQuery) SetTenantId(val tenant.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *EndorserUpdateQuery) SetMemberId(val member.MemberId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *EndorserUpdateQuery) SetGrade(val grade.Grade) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[2] = v
}

func (q *EndorserUpdateQuery) SetAvailableEndorsementCount(val endorserVal.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[3] = v
}

func (q *EndorserUpdateQuery) SetPendingEndorsementCount(val endorserVal.EndorsementCount) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[4] = v
}

func (q *EndorserUpdateQuery) SetVersion(val uint) {
	q.params[5] = val
}

func (q *EndorserUpdateQuery) SetCreatedAt(val time.Time) {
	q.params[6] = val
}

func (q *EndorserUpdateQuery) Evaluate(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	// TODO: Optimistic lock
	return s.Exec(q.sql(), q.params[:]...)
}
