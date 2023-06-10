package tenant

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewTenantRepository(session infrastructure.DbSession) TenantRepository {
	return TenantRepository{
		session: session,
	}
}

type TenantRepository struct {
	session infrastructure.DbSession
}

func (r TenantRepository) Insert(obj *tenant.Tenant) error {
	q := TenantInsertQuery{}
	obj.Export(&q)
	result, err := q.Execute(r.session)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows != 0 {
		return fmt.Errorf("wrong rows affected: %d", affectedRows)
	}
	return nil
}
