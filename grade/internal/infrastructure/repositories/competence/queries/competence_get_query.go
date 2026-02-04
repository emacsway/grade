package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

type CompetenceGetQuery struct {
	Id competenceVal.CompetenceId
}

func (q CompetenceGetQuery) sql() string {
	return `
		SELECT
		tenant_id, competence_id, name, owner_id, created_at, version
		FROM competence
		WHERE tenant_id=$1 AND competence_id=$2`
}

func (q CompetenceGetQuery) params() []any {
	var idExp competenceVal.CompetenceIdExporter
	q.Id.Export(&idExp)
	return []any{idExp.TenantId, idExp.CompetenceId}
}

func (q *CompetenceGetQuery) Get(s session.Session) (*competence.Competence, error) {
	rec := &competence.CompetenceReconstitutor{}
	err := s.(session.DbSession).Connection().QueryRow(q.sql(), q.params()...).Scan(
		&rec.Id.TenantId, &rec.Id.CompetenceId, &rec.Name, &rec.OwnerId.MemberId,
		&rec.CreatedAt, &rec.Version,
	)
	// TODO: Find the right place???
	rec.OwnerId.TenantId = rec.Id.TenantId
	if err != nil {
		return nil, err
	}
	return rec.Reconstitute()
}
