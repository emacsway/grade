package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
)

func NewMemberFaker(
	session infrastructure.DbSession,
	opts ...member.MemberFakerOption,
) *member.MemberFaker {
	opts = append(
		opts,
		member.WithTransientId(),
		member.WithRepository(NewMemberRepository(session)),
		// TODO: Is it an argument to pass a session to repository method?
		member.WithTenantFaker(tenantRepo.NewTenantFaker(session)),
	)
	return member.NewMemberFaker(opts...)
}
