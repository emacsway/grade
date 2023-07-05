package queries

import (
	"fmt"

	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type ArtifactPkGetQuery struct {
	Id tenantVal.TenantId
}

func (q ArtifactPkGetQuery) sql() string {
	return fmt.Sprintf(`SELECT  nextval('artifact_seq_%d')`, q.idValue())
}

func (q ArtifactPkGetQuery) idValue() uint {
	var id exporters.UintExporter
	q.Id.Export(&id)
	return uint(id)
}

func (q *ArtifactPkGetQuery) Get(s infrastructure.DbSessionSingleQuerier) (artifactVal.TenantArtifactId, error) {
	rec := &artifactVal.TenantArtifactIdReconstitutor{}
	rec.TenantId = q.idValue()
	err := s.QueryRow(q.sql()).Scan(&rec.ArtifactId)
	if err != nil {
		return artifactVal.TenantArtifactId{}, err
	}
	return rec.Reconstitute()
}
