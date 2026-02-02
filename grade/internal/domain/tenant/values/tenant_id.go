package values

import (
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/identity"
)

func NewTenantId(value uint) (TenantId, error) {
	id, err := identity.NewIntIdentity(value)
	if err != nil {
		return TenantId{}, err
	}
	return TenantId{&id}, nil
}

func NewTransientTenantId() TenantId {
	return TenantId{}
}

type TenantId struct {
	*identity.IntIdentity
}
