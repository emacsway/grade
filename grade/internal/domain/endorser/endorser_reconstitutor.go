package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

type EndorserReconstitutor struct {
	Id                        member.TenantMemberIdReconstitutor
	Grade                     uint8
	AvailableEndorsementCount uint
	PendingEndorsementCount   uint
	CreatedAt                 time.Time
	Version                   uint
}

func (r EndorserReconstitutor) Reconstitute() (*Endorser, error) {
	// Set here TenantId to other composite FK.
	id, err := r.Id.Reconstitute()
	if err != nil {
		return nil, err
	}
	g, err := grade.DefaultConstructor(r.Grade)
	if err != nil {
		return nil, err
	}
	availableEndorsementCount, err := NewEndorsementCount(r.AvailableEndorsementCount)
	if err != nil {
		return nil, err
	}
	pendingEndorsementCount, err := NewEndorsementCount(r.PendingEndorsementCount)
	if err != nil {
		return nil, err
	}
	return &Endorser{
		id:                        id,
		grade:                     g,
		availableEndorsementCount: availableEndorsementCount,
		pendingEndorsementCount:   pendingEndorsementCount,
		createdAt:                 r.CreatedAt,
		VersionedAggregate:        aggregate.NewVersionedAggregate(r.Version),
	}, nil
}
