package tenant

import (
	"time"
)

type TenantFakeFactoryOption func(*TenantFakeFactory)

func WithTransientId() TenantFakeFactoryOption {
	return func(f *TenantFakeFactory) {
		f.Id = 0
	}
}

func NewTenantFakeFactory(opts ...TenantFakeFactoryOption) TenantFakeFactory {
	f := TenantFakeFactory{
		Id:        TenantIdFakeValue,
		Name:      "Name1",
		CreatedAt: time.Now(),
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
	id, err := NewTenantId(f.Id)
	if err != nil {
		return nil, err
	}
	name, err := NewName(f.Name)
	if err != nil {
		return nil, err
	}
	return NewTenant(
		id, name, f.CreatedAt,
	)
}
