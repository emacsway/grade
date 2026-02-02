package competence

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/aggregate"
)

type CompetenceReconstitutor struct {
	Id        values.CompetenceIdReconstitutor
	Name      string
	OwnerId   member.MemberIdReconstitutor
	CreatedAt time.Time
	Version   uint
}

func (r CompetenceReconstitutor) Reconstitute() (*Competence, error) {
	id, err := r.Id.Reconstitute()
	if err != nil {
		return nil, err
	}
	name, err := values.NewName(r.Name)
	if err != nil {
		return nil, err
	}
	ownerId, err := r.OwnerId.Reconstitute()
	if err != nil {
		return nil, err
	}
	return &Competence{
		id:                 id,
		name:               name,
		ownerId:            ownerId,
		createdAt:          r.CreatedAt,
		VersionedAggregate: aggregate.NewVersionedAggregate(r.Version),
	}, nil
}
