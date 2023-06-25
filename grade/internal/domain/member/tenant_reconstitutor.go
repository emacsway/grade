package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

type MemberReconstitutor struct {
	TenantId  uint
	Id        uint
	Status    uint8
	FirstName string
	LastName  string
	CreatedAt time.Time
	Version   uint
}

func (r MemberReconstitutor) Reconstitute() (*Member, error) {
	id, err := values.NewTenantMemberId(r.TenantId, r.Id)
	if err != nil {
		return nil, err
	}
	status, err := values.NewStatus(r.Status)
	if err != nil {
		return nil, err
	}
	fullName, err := values.NewFullName(r.FirstName, r.LastName)
	if err != nil {
		return nil, err
	}
	return &Member{
		id:                 id,
		status:             status,
		fullName:           fullName,
		createdAt:          r.CreatedAt,
		VersionedAggregate: aggregate.NewVersionedAggregate(r.Version),
	}, nil
}
