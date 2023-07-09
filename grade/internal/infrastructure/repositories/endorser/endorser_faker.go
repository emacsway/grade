package endorser

import (
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
)

func NewEndorserFaker(
	session infrastructure.DbSession,
	opts ...endorser.EndorserFakerOption,
) *endorser.EndorserFaker {
	opts = append(
		opts,
		endorser.WithRepository(NewEndorserRepository(session)),
		endorser.WithMemberFaker(memberRepo.NewMemberFaker(session)),
	)
	return endorser.NewEndorserFaker(opts...)
}
