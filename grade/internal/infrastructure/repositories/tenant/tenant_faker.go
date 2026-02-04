package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantFaker(
	opts ...tenant.TenantFakerOption,
) *tenant.TenantFaker {
	opts = append(
		opts,
		tenant.WithTransientId(),
		tenant.WithRepository(NewTenantRepository()),
	)
	return tenant.NewTenantFaker(opts...)
}
