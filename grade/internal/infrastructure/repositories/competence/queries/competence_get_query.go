package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type CompetenceGetQuery struct {
	Id competenceVal.TenantCompetenceId
}

func (q CompetenceGetQuery) sql() string {
	return `
		SELECT
		tenant_id, competence_id, name, owner_id, created_at, version
		FROM competence
		WHERE tenant_id=$1 AND competence_id=$2`
}

func (q CompetenceGetQuery) params() []any {
	var idExp competenceVal.TenantCompetenceIdExporter
	q.Id.Export(&idExp)
	return []any{idExp.TenantId, idExp.CompetenceId}
}

func (q *CompetenceGetQuery) Get(s infrastructure.DbSessionSingleQuerier) (*competence.Competence, error) {
	rec := &competence.CompetenceReconstitutor{}
	err := s.QueryRow(q.sql(), q.params()...).Scan(
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
