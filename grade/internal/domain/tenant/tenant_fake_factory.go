package tenant

import (
	"time"
)

func NewTenantFakeFactory() TenantFakeFactory {
	return TenantFakeFactory{
		Id:        TenantIdFakeValue,
		Name:      "Name1",
		CreatedAt: time.Now(),
	}
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
