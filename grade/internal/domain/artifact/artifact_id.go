package artifact

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewArtifactId(value uint64) ArtifactId {
	return ArtifactId{seedwork.NewIdentity[uint64](value)}
}

type ArtifactId struct {
	seedwork.Identity[uint64]
}
