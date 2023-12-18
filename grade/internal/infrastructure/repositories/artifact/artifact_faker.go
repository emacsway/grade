package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	competenceRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/competence"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewArtifactFaker(
	session session.DbSession,
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
