package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewTenantId(value uuid.Uuid) (TenantId, error) {
	id, err := seedwork.NewUuidIdentity(value)
	if err != nil {
		return TenantId{}, err
	}
	return TenantId{id}, nil
}

type TenantId struct {
	seedwork.UuidIdentity
}
