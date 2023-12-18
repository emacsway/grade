package queries

import (
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
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
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *EndorserDeleteQuery) SetMemberId(val member.InternalMemberId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *EndorserDeleteQuery) Evaluate(s session.DbSessionExecutor) (session.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
