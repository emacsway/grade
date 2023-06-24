package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type TenantGetQuery struct {
	Id tenantVal.TenantId
}

func (q TenantGetQuery) sql() string {
	return `
		SELECT
		id, name, version, created_at
		FROM tenant
		WHERE id=$1`
}

func (q TenantGetQuery) params() []any {
	var id exporters.UintExporter
	q.Id.Export(&id)
	return []any{id}
}

func (q *TenantGetQuery) Get(s infrastructure.DbSessionSingleQuerier) (*tenant.Tenant, error) {
	rec := &tenant.TenantReconstitutor{}
	err := s.QueryRow(q.sql(), q.params()...).Scan(
		&rec.Id, &rec.Name, &rec.Version, &rec.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return rec.Reconstitute()
}
