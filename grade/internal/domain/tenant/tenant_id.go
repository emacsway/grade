package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewTenantId(value uuid.Uuid) (TenantId, error) {
	id, err := identity.NewUuidIdentity(value)
	if err != nil {
		return TenantId{}, err
	}
	return TenantId{id}, nil
}

type TenantId struct {
	identity.UuidIdentity
}
