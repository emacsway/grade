package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/aggregate"
)

type MemberReconstitutor struct {
	Id        values.MemberIdReconstitutor
	Status    uint8
	FullName  values.FullNameReconstitutor
	CreatedAt time.Time
	Version   uint
}

func (r MemberReconstitutor) Reconstitute() (*Member, error) {
	id, err := r.Id.Reconstitute()
	if err != nil {
		return nil, err
	}
	status, err := values.NewStatus(r.Status)
	if err != nil {
		return nil, err
	}
	fullName, err := r.FullName.Reconstitute()
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
