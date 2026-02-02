package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewCompetenceFaker(
	currentSession session.DbSession,
	opts ...competence.CompetenceFakerOption,
) *competence.CompetenceFaker {
	opts = append(
		opts,
		competence.WithTransientId(),
		competence.WithRepository(NewCompetenceRepository(currentSession)),
		competence.WithMemberFaker(memberRepo.NewMemberFaker(currentSession)),
	)
	return competence.NewCompetenceFaker(opts...)
}
