package endorser

import (
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewEndorserFaker(
	currentSession session.DbSession,
	opts ...endorser.EndorserFakerOption,
) *endorser.EndorserFaker {
	opts = append(
		opts,
		endorser.WithRepository(NewEndorserRepository(currentSession)),
		endorser.WithMemberFaker(memberRepo.NewMemberFaker(currentSession)),
	)
	return endorser.NewEndorserFaker(opts...)
}
