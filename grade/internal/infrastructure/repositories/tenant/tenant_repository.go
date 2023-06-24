package tenant

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/tenant"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewTenantRepository(session infrastructure.DbSession) *TenantRepository {
	return &TenantRepository{
		session: session,
	}
}

type TenantRepository struct {
	session infrastructure.DbSession
}

func (r TenantRepository) Insert(obj *tenant.Tenant) error {
	q := TenantInsertQuery{}
	obj.Export(&q)
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

func (r TenantRepository) Get(id tenantVal.TenantId) (*tenant.Tenant, error) {
	q := TenantGetQuery{id}
	return q.Get(r.session)
}
