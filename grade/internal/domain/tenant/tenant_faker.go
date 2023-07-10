package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/faker"
	"github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

type TenantFakerOption func(*TenantFaker)

func WithTransientId() TenantFakerOption {
	return func(f *TenantFaker) {
		f.Id = 0
	}
}

func WithRepository(repo TenantRepository) TenantFakerOption {
	return func(f *TenantFaker) {
		f.Repository = repo
	}
}

func NewTenantFaker(opts ...TenantFakerOption) *TenantFaker {
	aFaker := faker.NewFaker()
	f := &TenantFaker{
		Id:        values.TenantIdFakeValue,
		Name:      aFaker.Company(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
	repo := &TenantDummyRepository{
		TenantFaker: f,
	}
	f.Repository = repo
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type TenantFaker struct {
	Id         uint
	Name       string
	CreatedAt  time.Time
	Repository TenantRepository
}

func (f *TenantFaker) Create() (*Tenant, error) {
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

type TenantDummyRepository struct {
	TenantFaker *TenantFaker
}

func (r *TenantDummyRepository) Insert(agg *Tenant) error {
	// r.TenantFaker.Id += 1 // Do not do this, since TenantFaker.Id is an autoincrement PK accessor.
	// This should be exactly as agg.id
	// Also this value will be reseted by f.Id = uint(aggExp.Id)
	return nil
}

type TenantInsertDummyQuery struct {
	pkSetter func(any) error
}

func (q *TenantInsertDummyQuery) SetId(val values.TenantId) {
	q.pkSetter = val.Scan
}
func (q *TenantInsertDummyQuery) SetName(val values.Name)    {}
func (q *TenantInsertDummyQuery) SetCreatedAt(val time.Time) {}
func (q *TenantInsertDummyQuery) SetVersion(val uint)        {}
func (q *TenantInsertDummyQuery) Evaluate(f *TenantFaker) error {
	return q.pkSetter(f.Id + 1)
}
