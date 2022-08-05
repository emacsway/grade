package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewArtifactId(value uuid.Uuid) (ArtifactId, error) {
	id, err := identity.NewUuidIdentity(value)
	if err != nil {
		return ArtifactId{}, err
	}
	return ArtifactId{id}, nil
}

type ArtifactId struct {
	identity.UuidIdentity
}
