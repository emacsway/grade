package queries

import (
	"fmt"

	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

type ArtifactNextIdGetQuery struct {
	TenantId tenantVal.TenantId
}

func (q ArtifactNextIdGetQuery) sql() string {
	return fmt.Sprintf(`SELECT nextval('artifact_seq_%d'::regclass)`, q.tenantIdValue())
}

func (q ArtifactNextIdGetQuery) tenantIdValue() uint {
	var tenantIdExp uint
	q.TenantId.Export(func(v uint) { tenantIdExp = v })
	return tenantIdExp
}

func (q *ArtifactNextIdGetQuery) Get(s session.DbSessionSingleQuerier) (artifactVal.ArtifactId, error) {
	rec := &artifactVal.ArtifactIdReconstitutor{}
	rec.TenantId = q.tenantIdValue()
	err := s.QueryRow(q.sql()).Scan(&rec.ArtifactId)
	if err != nil {
		return artifactVal.ArtifactId{}, err
	}
	return rec.Reconstitute()
}
