package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

type MemberGetQuery struct {
	Id memberVal.MemberId
}

func (q MemberGetQuery) sql() string {
	return `
		SELECT
		tenant_id, member_id, status, first_name, last_name, created_at, version
		FROM member
		WHERE tenant_id=$1 AND member_id=$2`
}

func (q MemberGetQuery) params() []any {
	var idExp memberVal.MemberIdExporter
	q.Id.Export(&idExp)
	return []any{idExp.TenantId, idExp.MemberId}
}

func (q *MemberGetQuery) Get(s session.DbSessionSingleQuerier) (*member.Member, error) {
	rec := &member.MemberReconstitutor{}
	err := s.QueryRow(q.sql(), q.params()...).Scan(
		&rec.Id.TenantId, &rec.Id.MemberId, &rec.Status, &rec.FullName.FirstName, &rec.FullName.LastName,
		&rec.CreatedAt, &rec.Version,
	)
	if err != nil {
		return nil, err
	}
	return rec.Reconstitute()
}
