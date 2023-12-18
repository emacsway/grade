package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewCompetenceFaker(
	session session.DbSession,
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
