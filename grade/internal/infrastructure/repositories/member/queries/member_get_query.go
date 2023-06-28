package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type MemberGetQuery struct {
	Id memberVal.TenantMemberId
}

func (q MemberGetQuery) sql() string {
	return `
		SELECT
		tenant_id, member_id, status, first_name, last_name, version, created_at
		FROM member
		WHERE tenant_id=$1 AND member_id=$2`
}

func (q MemberGetQuery) params() []any {
	var idExp memberVal.TenantMemberIdExporter
	q.Id.Export(&idExp)
	return []any{idExp.TenantId, idExp.MemberId}
}

func (q *MemberGetQuery) Get(s infrastructure.DbSessionSingleQuerier) (*member.Member, error) {
	rec := &member.MemberReconstitutor{}
	err := s.QueryRow(q.sql(), q.params()...).Scan(
		&rec.Id.TenantId, &rec.Id.MemberId, &rec.Status, &rec.FirstName, &rec.LastName, &rec.Version, &rec.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return rec.Reconstitute()
}
