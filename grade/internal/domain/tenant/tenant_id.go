package tenant

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewTenantId(value uint) (TenantId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return TenantId{}, err
	}
	return TenantId{id}, nil
}

type TenantId struct {
	identity.IntIdentity
}
