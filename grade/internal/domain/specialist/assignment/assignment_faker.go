package assignment

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment/values"
)

func NewAssignmentFaker() AssignmentFaker {
	return AssignmentFaker{
		SpecialistId:      member.NewMemberIdFaker(),
		SpecialistVersion: 2,
		AssignedGrade:     1,
		Reason:            "Any",
		CreatedAt:         time.Now().Truncate(time.Microsecond),
	}
}

type AssignmentFaker struct {
	SpecialistId      member.MemberIdFaker
	SpecialistVersion uint
	AssignedGrade     uint8
	Reason            string
	CreatedAt         time.Time
}

func (f AssignmentFaker) Create() (Assignment, error) {
	specialistId, err := f.SpecialistId.Create()
	if err != nil {
		return Assignment{}, err
	}
	assignedGrade, err := grade.DefaultConstructor(f.AssignedGrade)
	if err != nil {
		return Assignment{}, err
	}
	reason, err := values.NewReason(f.Reason)
	if err != nil {
		return Assignment{}, err
	}
	return NewAssignment(specialistId, f.SpecialistVersion, assignedGrade, reason, f.CreatedAt)
}
