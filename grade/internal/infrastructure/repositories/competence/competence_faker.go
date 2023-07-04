package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewCompetenceFaker(
	session infrastructure.DbSession,
	opts ...competence.CompetenceFakerOption,
) *competence.CompetenceFaker {
	opts = append(
		opts,
		competence.WithTransientId(),
		competence.WithRepository(NewCompetenceRepository(session)),
	)
	return competence.NewCompetenceFaker(opts...)
}
