package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewMemberFaker(
	currentSession session.DbSession,
	opts ...member.MemberFakerOption,
) *member.MemberFaker {
	opts = append(
		opts,
		member.WithTransientId(),
		member.WithRepository(NewMemberRepository(currentSession)),
		// TODO: Is this a reason to pass a session to repository method?
		member.WithTenantFaker(tenantRepo.NewTenantFaker(currentSession)),
	)
	return member.NewMemberFaker(opts...)
}
