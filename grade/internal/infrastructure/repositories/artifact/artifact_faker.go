package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewArtifactFaker(
	session infrastructure.DbSession,
	opts ...artifact.ArtifactFakerOption,
) *artifact.ArtifactFaker {
	opts = append(
		opts,
		artifact.WithTransientId(),
		artifact.WithRepository(NewArtifactRepository(session)),
	)
	return artifact.NewArtifactFaker(opts...)
}
