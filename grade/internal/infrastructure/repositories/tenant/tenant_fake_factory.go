package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewTenantFakeFactory(
	session infrastructure.DbSession,
	opts ...tenant.TenantFakeFactoryOption,
) *tenant.TenantFakeFactory {
	opts = append(
		opts,
		tenant.WithTransientId(),
		tenant.WithRepository(NewTenantRepository(session)),
	)
	return tenant.NewTenantFakeFactory(opts...)
}
