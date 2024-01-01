package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewTenantFaker(
	currentSession session.DbSession,
	opts ...tenant.TenantFakerOption,
) *tenant.TenantFaker {
	opts = append(
		opts,
		tenant.WithTransientId(),
		tenant.WithRepository(NewTenantRepository(currentSession)),
	)
	return tenant.NewTenantFaker(opts...)
}
