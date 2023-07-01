package endorser

import (
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewEndorserFaker(
	session infrastructure.DbSession,
	opts ...endorser.EndorserFakerOption,
) *endorser.EndorserFaker {
	opts = append(
		opts,
		endorser.WithRepository(NewEndorserRepository(session)),
	)
	return endorser.NewEndorserFaker(opts...)
}
