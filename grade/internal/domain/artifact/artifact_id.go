package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewArtifactId(value uuid.Uuid) (ArtifactId, error) {
	id, err := seedwork.NewUuidIdentity(value)
	if err != nil {
		return ArtifactId{}, err
	}
	return ArtifactId{id}, nil
}

type ArtifactId struct {
	seedwork.UuidIdentity
}
