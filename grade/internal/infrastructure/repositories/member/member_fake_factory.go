package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewMemberFakeFactory(
	session infrastructure.DbSession,
	opts ...member.MemberFakeFactoryOption,
) *member.MemberFakeFactory {
	opts = append(
		opts,
		member.WithTransientId(),
		member.WithRepository(NewMemberRepository(session)),
	)
	return member.NewMemberFakeFactory(opts...)
}
