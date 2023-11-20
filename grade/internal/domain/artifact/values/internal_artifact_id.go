package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewInternalArtifactId(value uint) (InternalArtifactId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return InternalArtifactId{}, err
	}
	return InternalArtifactId{id}, nil
}

type InternalArtifactId struct {
	identity.IntIdentity
}
