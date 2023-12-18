package endorser

import (
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewEndorserFaker(
	session session.DbSession,
	opts ...endorser.EndorserFakerOption,
) *endorser.EndorserFaker {
	opts = append(
		opts,
		endorser.WithRepository(NewEndorserRepository(session)),
		endorser.WithMemberFaker(memberRepo.NewMemberFaker(session)),
	)
	return endorser.NewEndorserFaker(opts...)
}
