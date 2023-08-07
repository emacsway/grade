package queries

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/seedwork/repository"
)

type ArtifactGetQuery struct {
	repository.EventGetQuery
}

func (q *ArtifactGetQuery) Get(s infrastructure.DbSessionQuerier) (*artifact.Artifact, error) {
	stream, err := q.Stream(s)
	if err != nil {
		return nil, err
	}
	rec := &artifact.ArtifactReconstitutor{
		PastEvents: stream,
	}
	return rec.Reconstitute()
}
