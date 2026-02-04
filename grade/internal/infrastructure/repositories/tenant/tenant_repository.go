package tenant

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/tenant"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant/queries"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewTenantRepository() *TenantRepository {
	return &TenantRepository{}
}

type TenantRepository struct{}

func (r *TenantRepository) Insert(s session.Session, agg *tenant.Tenant) error {
	q := &queries.TenantInsertQuery{}
	agg.Export(q)
	result, err := q.Evaluate(s)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	if lastInsertId == 0 {
		return fmt.Errorf("wrong LastInsertId: %d", lastInsertId)
	}
	return nil
}

func (r *TenantRepository) Get(s session.Session, id tenantVal.TenantId) (*tenant.Tenant, error) {
	q := queries.TenantGetQuery{Id: id}
	return q.Get(s)
}
