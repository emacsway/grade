package queries

import (
	"fmt"

	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type ArtifactNextIdGetQuery struct {
	TenantId tenantVal.TenantId
}

func (q ArtifactNextIdGetQuery) sql() string {
	return fmt.Sprintf(`SELECT nextval('artifact_seq_%d'::regclass)`, q.tenantIdValue())
}

func (q ArtifactNextIdGetQuery) tenantIdValue() uint {
	var tenantIdExp exporters.UintExporter
	q.TenantId.Export(&tenantIdExp)
	return uint(tenantIdExp)
}

func (q *ArtifactNextIdGetQuery) Get(s infrastructure.DbSessionSingleQuerier) (artifactVal.ArtifactId, error) {
	rec := &artifactVal.ArtifactIdReconstitutor{}
	rec.TenantId = q.tenantIdValue()
	err := s.QueryRow(q.sql()).Scan(&rec.ArtifactId)
	if err != nil {
		return artifactVal.ArtifactId{}, err
	}
	return rec.Reconstitute()
}
