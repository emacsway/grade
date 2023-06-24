package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/faker"
	"github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

type TenantFakeFactoryOption func(*TenantFakeFactory)

func WithTransientId() TenantFakeFactoryOption {
	return func(f *TenantFakeFactory) {
		f.Id = 0
	}
}

func NewTenantFakeFactory(opts ...TenantFakeFactoryOption) TenantFakeFactory {
	aFaker := faker.NewFaker()
	f := TenantFakeFactory{
		Id:        values.TenantIdFakeValue,
		Name:      aFaker.Company(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
	for _, opt := range opts {
		opt(&f)
	}
	return f
}

type TenantFakeFactory struct {
	Id        uint
	Name      string
	CreatedAt time.Time
}

func (f TenantFakeFactory) Create() (*Tenant, error) {
	id, err := values.NewTenantId(f.Id)
	if err != nil {
		return nil, err
	}
	name, err := values.NewName(f.Name)
	if err != nil {
		return nil, err
	}
	return NewTenant(
		id, name, f.CreatedAt,
	)
}
