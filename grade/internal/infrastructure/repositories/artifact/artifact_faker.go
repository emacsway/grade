package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	competenceRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/competence"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
)

func NewArtifactFaker(
	opts ...artifact.ArtifactFakerOption,
) *artifact.ArtifactFaker {
	opts = append(
		opts,
		artifact.WithTransientId(),
		artifact.WithRepository(NewArtifactRepository()),
		artifact.WithMemberFaker(memberRepo.NewMemberFaker()),
		artifact.WithCompetenceFaker(competenceRepo.NewCompetenceFaker()),
	)
	return artifact.NewArtifactFaker(opts...)
}
