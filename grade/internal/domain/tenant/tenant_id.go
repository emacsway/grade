package tenant

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewTenantId(value uint64) (TenantId, error) {
	id, err := seedwork.NewIdentity[uint64](value)
	if err != nil {
		return TenantId{}, err
	}
	return TenantId{id}, nil
}

type TenantId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64], interfaces.Exporter[uint64]]
}
