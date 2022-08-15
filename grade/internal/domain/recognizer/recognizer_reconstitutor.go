package recognizer

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

type RecognizerReconstitutor struct {
	Id                        member.TenantMemberIdReconstitutor
	Grade                     uint8
	AvailableEndorsementCount uint
	PendingEndorsementCount   uint
	CreatedAt                 time.Time
	Version                   uint
}

func (r RecognizerReconstitutor) Reconstitute() (*Recognizer, error) {
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
	return &Recognizer{
		id:                        id,
		grade:                     g,
		availableEndorsementCount: availableEndorsementCount,
		pendingEndorsementCount:   pendingEndorsementCount,
		createdAt:                 r.CreatedAt,
		VersionedAggregate:        aggregate.NewVersionedAggregate(r.Version),
	}, nil
}
