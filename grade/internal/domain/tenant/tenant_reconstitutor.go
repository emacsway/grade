package tenant

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

type TenantReconstitutor struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	Version   uint
}

func (r TenantReconstitutor) Reconstitute() (*Tenant, error) {
	id, err := NewTenantId(r.Id)
	if err != nil {
		return nil, err
	}
	name, err := NewName(r.Name)
	if err != nil {
		return nil, err
	}
	return &Tenant{
		id:                 id,
		name:               name,
		createdAt:          r.CreatedAt,
		VersionedAggregate: aggregate.NewVersionedAggregate(r.Version),
	}, nil
}
