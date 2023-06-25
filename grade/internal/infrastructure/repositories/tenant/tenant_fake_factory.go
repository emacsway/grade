package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewTenantFakeFactory(
	session infrastructure.DbSession,
	opts ...tenant.TenantFakeFactoryOption,
) *TenantFakeFactory {
	opts = append(opts, tenant.WithTransientId())
	return &TenantFakeFactory{
		tenant.NewTenantFakeFactory(opts...),
		NewTenantRepository(session),
	}
}

type TenantFakeFactory struct {
	tenant.TenantFakeFactory
	Repository *TenantRepository
}

func (f *TenantFakeFactory) Create() (*tenant.Tenant, error) {
	var aggExp tenant.TenantExporter
	agg, err := f.TenantFakeFactory.Create()
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id = uint(aggExp.Id)
	return agg, nil
}
