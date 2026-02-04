package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
)

func NewMemberFaker(
	opts ...member.MemberFakerOption,
) *member.MemberFaker {
	opts = append(
		opts,
		member.WithTransientId(),
		member.WithRepository(NewMemberRepository()),
		member.WithTenantFaker(tenantRepo.NewTenantFaker()),
	)
	return member.NewMemberFaker(opts...)
}
