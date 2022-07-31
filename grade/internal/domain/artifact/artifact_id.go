package artifact

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewArtifactId(value uint64) (ArtifactId, error) {
	id, err := seedwork.NewUint64Identity(value)
	if err != nil {
		return ArtifactId{}, err
	}
	return ArtifactId{id}, nil
}

type ArtifactId struct {
	seedwork.Uint64Identity
}
