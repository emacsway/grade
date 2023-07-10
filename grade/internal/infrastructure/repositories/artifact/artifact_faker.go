package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	competenceRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/competence"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
)

func NewArtifactFaker(
	session infrastructure.DbSession,
	opts ...artifact.ArtifactFakerOption,
) *artifact.ArtifactFaker {
	opts = append(
		opts,
		artifact.WithTransientId(),
		artifact.WithRepository(NewArtifactRepository(session)),
		artifact.WithMemberFaker(memberRepo.NewMemberFaker(session)),
		artifact.WithCompetenceFaker(competenceRepo.NewCompetenceFaker(session)),
	)
	return artifact.NewArtifactFaker(opts...)
}
