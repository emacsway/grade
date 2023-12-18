package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

type EndorserGetQuery struct {
	Id memberVal.MemberId
}

func (q EndorserGetQuery) sql() string {
	return `
		SELECT
		tenant_id, member_id, grade, available_endorsement_count,
		pending_endorsement_count, created_at, version
		FROM endorser
		WHERE tenant_id=$1 AND member_id=$2`
}

func (q EndorserGetQuery) params() []any {
	var idExp memberVal.MemberIdExporter
	q.Id.Export(&idExp)
	return []any{idExp.TenantId, idExp.MemberId}
}

func (q *EndorserGetQuery) Get(s session.DbSessionSingleQuerier) (*endorser.Endorser, error) {
	rec := &endorser.EndorserReconstitutor{}
	err := s.QueryRow(q.sql(), q.params()...).Scan(
		&rec.Id.TenantId, &rec.Id.MemberId, &rec.Grade, &rec.AvailableEndorsementCount,
		&rec.PendingEndorsementCount, &rec.CreatedAt, &rec.Version,
	)
	if err != nil {
		return nil, err
	}
	return rec.Reconstitute()
}
