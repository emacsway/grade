package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/faker"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
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
	f := &TenantFaker{
		Id: values.TenantIdFakeValue,
	}
	f.fake()
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
	agg        *Tenant
}

func (f *TenantFaker) fake() {
	aFaker := faker.NewFaker()
	f.Name = aFaker.Company()
	f.CreatedAt = time.Now().Truncate(time.Microsecond)
}

func (f *TenantFaker) Next() error {
	f.fake()
	f.Id += 1
	f.agg = nil
	return nil
}

func (f *TenantFaker) Create(s session.Session) (*Tenant, error) {
	var aggExp TenantExporter
	if f.agg != nil {
		return f.agg, nil
	}
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
	err = f.Repository.Insert(s, agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id = uint(aggExp.Id)
	f.agg = agg
	return agg, nil
}

// unidirectional flow of changes
func (f *TenantFaker) SetId(val uint) {
	f.Id = val
}

func (f *TenantFaker) BuildDependencies(s session.Session) (err error) {
	return nil
}

type TenantRepository interface {
	Insert(session.Session, *Tenant) error
}

type TenantDummyRepository struct {
	TenantFaker *TenantFaker
}

func (r *TenantDummyRepository) Insert(s session.Session, agg *Tenant) error {
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
