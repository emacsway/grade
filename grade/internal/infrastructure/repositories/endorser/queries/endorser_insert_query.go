package queries

import (
	"time"

	endorserVal "github.com/emacsway/grade/grade/internal/domain/endorser/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

type EndorserInsertQuery struct {
	params [7]any
}

func (q EndorserInsertQuery) sql() string {
	return `
		INSERT INTO endorser
		(tenant_id, member_id, grade, available_endorsement_count, pending_endorsement_count, created_at, version)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)`
}

func (q *EndorserInsertQuery) SetId(val memberVal.MemberId) {
	val.Export(q)
}

func (q *EndorserInsertQuery) SetTenantId(val tenantVal.TenantId) {
	val.Export(func(v uint) { q.params[0] = v })
}

func (q *EndorserInsertQuery) SetMemberId(val memberVal.InternalMemberId) {
	val.Export(func(v uint) { q.params[1] = v })
}

func (q *EndorserInsertQuery) SetGrade(val grade.Grade) {
	val.Export(func(v uint8) { q.params[2] = v })
}

func (q *EndorserInsertQuery) SetAvailableEndorsementCount(val endorserVal.EndorsementCount) {
	val.Export(func(v uint) { q.params[3] = v })
}

func (q *EndorserInsertQuery) SetPendingEndorsementCount(val endorserVal.EndorsementCount) {
	val.Export(func(v uint) { q.params[4] = v })
}

func (q *EndorserInsertQuery) SetCreatedAt(val time.Time) {
	q.params[5] = val
}

func (q *EndorserInsertQuery) SetVersion(val uint) {
	q.params[6] = val
}

func (q *EndorserInsertQuery) Evaluate(s session.DbSessionExecutor) (session.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
