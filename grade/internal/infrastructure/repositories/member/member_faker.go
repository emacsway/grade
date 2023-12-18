package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewMemberFaker(
	session session.DbSession,
	opts ...member.MemberFakerOption,
) *member.MemberFaker {
	opts = append(
		opts,
		member.WithTransientId(),
		member.WithRepository(NewMemberRepository(session)),
		// TODO: Is this a reason to pass a session to repository method?
		member.WithTenantFaker(tenantRepo.NewTenantFaker(session)),
	)
	return member.NewMemberFaker(opts...)
}
