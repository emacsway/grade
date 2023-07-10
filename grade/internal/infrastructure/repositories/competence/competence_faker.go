package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
)

func NewCompetenceFaker(
	session infrastructure.DbSession,
	opts ...competence.CompetenceFakerOption,
) *competence.CompetenceFaker {
	opts = append(
		opts,
		competence.WithTransientId(),
		competence.WithRepository(NewCompetenceRepository(session)),
		competence.WithMemberFaker(memberRepo.NewMemberFaker(session)),
	)
	return competence.NewCompetenceFaker(opts...)
}
