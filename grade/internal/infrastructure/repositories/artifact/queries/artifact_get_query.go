package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type ArtifactGetQuery struct {
	Id         values.TenantArtifactId
	streamType string
}

func (q ArtifactGetQuery) sql() string {
	return `
		SELECT
		    tenant_id, stream_type, stream_id, stream_position, event_type, event_version, payload, metadata
		FROM
			event_log
		WHERE
			tenant_id=$1 AND stream_type=$2 AND stream_id=$3
		ORDER BY
			tenant_id, stream_type, stream_id, stream_position`
}
func (q ArtifactGetQuery) params() []any {
	var idExp values.TenantArtifactIdExporter
	q.Id.Export(&idExp)
	return []any{idExp.TenantId, q.streamType, idExp.ArtifactId.String()}
}
func (q *ArtifactGetQuery) SetStreamType(val string) {
	q.streamType = val
}
func (q *ArtifactGetQuery) Get(s infrastructure.DbSessionQuerier) (*artifact.Artifact, error) {
	rec := &artifact.ArtifactReconstitutor{}
	rows, err := s.Query(q.sql(), q.params()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan() // TODO: implement me
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return rec.Reconstitute()
}
