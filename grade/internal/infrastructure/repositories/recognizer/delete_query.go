package recognizer

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

type RecognizerDeleteQuery struct {
	params [2]any
}

func (q RecognizerDeleteQuery) sql() string {
	return `
		DELETE FROM recognizer
		WHERE tenant_id = $1, member_id=$2`
}

func (q *RecognizerDeleteQuery) SetTenantId(val tenant.TenantId) {
	var v exporters.UuidExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *RecognizerDeleteQuery) SetMemberId(val member.MemberId) {
	var v exporters.UuidExporter
	val.Export(&v)
	q.params[1] = v
}
