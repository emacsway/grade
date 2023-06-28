package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewTenantFaker(
	session infrastructure.DbSession,
	opts ...tenant.TenantFakerOption,
) *tenant.TenantFaker {
	opts = append(
		opts,
		tenant.WithTransientId(),
		tenant.WithRepository(NewTenantRepository(session)),
	)
	return tenant.NewTenantFaker(opts...)
}
