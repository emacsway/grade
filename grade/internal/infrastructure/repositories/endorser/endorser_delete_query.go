package endorser

import (
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
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

func (q *EndorserDeleteQuery) SetMemberId(val member.MemberId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[1] = v
}

func (q *EndorserDeleteQuery) Evaluate(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	return s.Exec(q.sql(), q.params[:]...)
}
