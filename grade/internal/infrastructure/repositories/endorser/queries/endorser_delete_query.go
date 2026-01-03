package queries

import (
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

type EndorserDeleteQuery struct {
	params [2]any
}

func (q EndorserDeleteQuery) sql() string {
	return `
		DELETE FROM endorser
		WHERE tenant_id = $1 AND member_id=$2`
}

func (q *EndorserDeleteQuery) SetTenantId(val tenant.TenantId) {
	val.Export(func(v uint) { q.params[0] = v })
}

func (q *EndorserDeleteQuery) SetMemberId(val member.InternalMemberId) {
	val.Export(func(v uint) { q.params[1] = v })
}

func (q *EndorserDeleteQuery) Evaluate(s session.DbSessionExecutor) (session.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
