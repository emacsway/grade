package tenant

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/tenant"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant/queries"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewTenantRepository(currentSession session.DbSession) *TenantRepository {
	return &TenantRepository{
		session: currentSession,
	}
}

type TenantRepository struct {
	session session.DbSession
}

func (r *TenantRepository) Insert(agg *tenant.Tenant) error {
	q := &queries.TenantInsertQuery{}
	agg.Export(q)
	result, err := q.Evaluate(r.session)
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

func (r *TenantRepository) Get(id tenantVal.TenantId) (*tenant.Tenant, error) {
	q := queries.TenantGetQuery{Id: id}
	return q.Get(r.session)
}
