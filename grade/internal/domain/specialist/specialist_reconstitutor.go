package specialist

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/grade/grade/internal/domain/specialist/endorsement"
)

type SpecialistReconstitutor struct {
	Id                   member.TenantMemberIdReconstitutor
	Grade                uint8
	receivedEndorsements []endorsement.EndorsementReconstitutor
	assignments          []assignment.AssignmentReconstitutor
	CreatedAt            time.Time
	Version              uint
}

func (r SpecialistReconstitutor) Reconstitute() (*Specialist, error) {
	id, err := r.Id.Reconstitute()
	if err != nil {
		return nil, err
	}
	aGrade, err := grade.DefaultConstructor(r.Grade)
	if err != nil {
		return nil, err
	}
	receivedEndorsements := []endorsement.Endorsement{}
	for i := range r.receivedEndorsements {
		receivedEndorsement, err := r.receivedEndorsements[i].Reconstitute()
		if err != nil {
			return nil, err
		}
		receivedEndorsements = append(receivedEndorsements, receivedEndorsement)
	}

	assignments := []assignment.Assignment{}
	for i := range r.assignments {
		anAssignment, err := r.assignments[i].Reconstitute()
		if err != nil {
			return nil, err
		}
		assignments = append(assignments, anAssignment)
	}
	return &Specialist{
		id:                   id,
		grade:                aGrade,
		receivedEndorsements: receivedEndorsements,
		assignments:          assignments,
		createdAt:            r.CreatedAt,
		VersionedAggregate:   aggregate.NewVersionedAggregate(r.Version),
	}, nil
}
