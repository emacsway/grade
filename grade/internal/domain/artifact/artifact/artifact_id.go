package artifact

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewArtifactId(value uint64) ArtifactId {
	return ArtifactId{seedwork.NewIdentity[uint64](value)}
}

type ArtifactId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64]]
}
