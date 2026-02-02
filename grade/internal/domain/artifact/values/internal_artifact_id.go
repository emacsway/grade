package values

import (
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/identity"
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
