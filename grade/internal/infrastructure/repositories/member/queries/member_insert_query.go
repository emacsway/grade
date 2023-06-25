package queries

import (
	"time"

	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type MemberInsertQuery struct {
	params   [6]any
	pkSetter func(any) error
}

func (q MemberInsertQuery) sql() string {
	return `
		INSERT INTO member
		(tenant_id, status, first_name, last_name, version, created_at)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id`
}

func (q *MemberInsertQuery) SetId(val memberVal.TenantMemberId) {
	val.Export(q)
}

func (q *MemberInsertQuery) SetTenantId(val tenant.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}

func (q *MemberInsertQuery) SetMemberId(val memberVal.MemberId) {
	q.pkSetter = val.Scan
}

func (q *MemberInsertQuery) SetStatus(val memberVal.Status) {
	var v exporters.Uint8Exporter
	val.Export(&v)
	q.params[1] = v
}

func (q *MemberInsertQuery) SetFullName(val memberVal.FullName) {
	val.Export(q)
}

func (q *MemberInsertQuery) SetFirstName(val memberVal.FirstName) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[2] = v
}

func (q *MemberInsertQuery) SetLastName(val memberVal.LastName) {
	var v exporters.StringExporter
	val.Export(&v)
	q.params[3] = v
}

func (q *MemberInsertQuery) SetVersion(val uint) {
	q.params[1] = val
}

func (q *MemberInsertQuery) SetCreatedAt(val time.Time) {
	q.params[2] = val
}

func (q *MemberInsertQuery) Evaluate(s infrastructure.DbSessionExecutor) (infrastructure.Result, error) {
	result, err := s.Exec(q.sql(), q.params[:]...)
	if err != nil {
		return result, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return result, err
	}
	err = q.pkSetter(id)
	return result, err
}
