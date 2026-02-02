package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	competenceRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/competence"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewArtifactFaker(
	currentSession session.DbSession,
	opts ...artifact.ArtifactFakerOption,
) *artifact.ArtifactFaker {
	opts = append(
		opts,
		artifact.WithTransientId(),
		artifact.WithRepository(NewArtifactRepository(currentSession)),
		artifact.WithMemberFaker(memberRepo.NewMemberFaker(currentSession)),
		artifact.WithCompetenceFaker(competenceRepo.NewCompetenceFaker(currentSession)),
	)
	return artifact.NewArtifactFaker(opts...)
}
