package tenant

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewTenantId(value uint64) (TenantId, error) {
	id, err := seedwork.NewUint64Identity(value)
	if err != nil {
		return TenantId{}, err
	}
	return TenantId{id}, nil
}

type TenantId struct {
	seedwork.Uint64Identity
}
