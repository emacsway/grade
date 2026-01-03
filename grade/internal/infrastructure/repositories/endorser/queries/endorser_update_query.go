package queries

import (
	"time"

	endorserVal "github.com/emacsway/grade/grade/internal/domain/endorser/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
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

func (q *EndorserUpdateQuery) SetId(val member.MemberId) {
	val.Export(q)
}

func (q *EndorserUpdateQuery) SetTenantId(val tenant.TenantId) {
	val.Export(func(v uint) { q.params[0] = v })
}

func (q *EndorserUpdateQuery) SetMemberId(val member.InternalMemberId) {
	val.Export(func(v uint) { q.params[1] = v })
}

func (q *EndorserUpdateQuery) SetGrade(val grade.Grade) {
	val.Export(func(v uint8) { q.params[2] = v })
}

func (q *EndorserUpdateQuery) SetAvailableEndorsementCount(val endorserVal.EndorsementCount) {
	val.Export(func(v uint) { q.params[3] = v })
}

func (q *EndorserUpdateQuery) SetPendingEndorsementCount(val endorserVal.EndorsementCount) {
	val.Export(func(v uint) { q.params[4] = v })
}

func (q *EndorserUpdateQuery) SetVersion(val uint) {
	q.params[5] = val
}

func (q *EndorserUpdateQuery) SetCreatedAt(val time.Time) {
	q.params[6] = val
}

func (q *EndorserUpdateQuery) Evaluate(s session.DbSessionExecutor) (session.Result, error) {
	// TODO: Optimistic lock
	return s.Exec(q.sql(), q.params[:]...)
}
