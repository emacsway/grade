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

func WithRepository(repo TenantRepository) TenantFakeFactoryOption {
	return func(f *TenantFakeFactory) {
		f.Repository = repo
	}
}

func NewTenantFakeFactory(opts ...TenantFakeFactoryOption) *TenantFakeFactory {
	aFaker := faker.NewFaker()
	f := &TenantFakeFactory{
		Id:         values.TenantIdFakeValue,
		Name:       aFaker.Company(),
		CreatedAt:  time.Now().Truncate(time.Microsecond),
		Repository: TenantDummyRepository{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type TenantFakeFactory struct {
	Id         uint
	Name       string
	CreatedAt  time.Time
	Repository TenantRepository
}

func (f *TenantFakeFactory) Create() (*Tenant, error) {
	var aggExp TenantExporter
	id, err := values.NewTenantId(f.Id)
	if err != nil {
		return nil, err
	}
	name, err := values.NewName(f.Name)
	if err != nil {
		return nil, err
	}
	agg, err := NewTenant(
		id, name, f.CreatedAt,
	)
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

type TenantRepository interface {
	Insert(*Tenant) error
}

type TenantDummyRepository struct{}

func (r TenantDummyRepository) Insert(agg *Tenant) error {
	return nil
}
