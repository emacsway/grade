package queries

import (
	"time"

	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

type MemberInsertQuery struct {
	params   [6]any
	pkSetter func(any) error
}

func (q MemberInsertQuery) sql() string {
	return `
		INSERT INTO member
		(tenant_id, status, first_name, last_name, created_at, version)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING member_id`
}

func (q *MemberInsertQuery) SetId(val memberVal.MemberId) {
	val.Export(q)
}

func (q *MemberInsertQuery) SetTenantId(val tenant.TenantId) {
	val.Export(func(v uint) { q.params[0] = v })
}

func (q *MemberInsertQuery) SetMemberId(val memberVal.InternalMemberId) {
	q.pkSetter = val.Scan
}

func (q *MemberInsertQuery) SetStatus(val memberVal.Status) {
	val.Export(func(v uint8) { q.params[1] = v })
}

func (q *MemberInsertQuery) SetFullName(val memberVal.FullName) {
	val.Export(q)
}

func (q *MemberInsertQuery) SetFirstName(val memberVal.FirstName) {
	val.Export(func(v string) { q.params[2] = v })
}

func (q *MemberInsertQuery) SetLastName(val memberVal.LastName) {
	val.Export(func(v string) { q.params[3] = v })
}

func (q *MemberInsertQuery) SetCreatedAt(val time.Time) {
	q.params[4] = val
}

func (q *MemberInsertQuery) SetVersion(val uint) {
	q.params[5] = val
}

func (q *MemberInsertQuery) Evaluate(s session.Session) (session.Result, error) {
	result, err := s.(session.DbSession).Connection().Exec(q.sql(), q.params[:]...)
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
