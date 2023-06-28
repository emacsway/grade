package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewMemberFaker(
	session infrastructure.DbSession,
	opts ...member.MemberFakerOption,
) *member.MemberFaker {
	opts = append(
		opts,
		member.WithTransientId(),
		member.WithRepository(NewMemberRepository(session)),
	)
	return member.NewMemberFaker(opts...)
}
