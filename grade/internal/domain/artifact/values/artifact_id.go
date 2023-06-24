package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewArtifactId(value uint) (ArtifactId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return ArtifactId{}, err
	}
	return ArtifactId{id}, nil
}

type ArtifactId struct {
	identity.IntIdentity
}
