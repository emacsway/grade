package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
)

func NewCompetenceFaker(
	opts ...competence.CompetenceFakerOption,
) *competence.CompetenceFaker {
	opts = append(
		opts,
		competence.WithTransientId(),
		competence.WithRepository(NewCompetenceRepository()),
		competence.WithMemberFaker(memberRepo.NewMemberFaker()),
	)
	return competence.NewCompetenceFaker(opts...)
}
