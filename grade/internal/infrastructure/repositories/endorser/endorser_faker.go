package endorser

import (
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
)

func NewEndorserFaker(
	opts ...endorser.EndorserFakerOption,
) *endorser.EndorserFaker {
	opts = append(
		opts,
		endorser.WithRepository(NewEndorserRepository()),
		endorser.WithMemberFaker(memberRepo.NewMemberFaker()),
	)
	return endorser.NewEndorserFaker(opts...)
}
