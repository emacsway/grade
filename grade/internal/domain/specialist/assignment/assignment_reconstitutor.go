package assignment

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment/values"
)

type AssignmentReconstitutor struct {
	SpecialistId      member.TenantMemberIdReconstitutor
	SpecialistVersion uint
	AssignedGrade     uint8
	Reason            string
	CreatedAt         time.Time
}

func (r AssignmentReconstitutor) Reconstitute() (Assignment, error) {
	specialistId, err := r.SpecialistId.Reconstitute()
	if err != nil {
		return Assignment{}, err
	}
	assignedGrade, err := grade.DefaultConstructor(r.AssignedGrade)
	if err != nil {
		return Assignment{}, err
	}
	reason, err := values.NewReason(r.Reason)
	if err != nil {
		return Assignment{}, err
	}
	return Assignment{
		specialistId:      specialistId,
		specialistVersion: r.SpecialistVersion,
		assignedGrade:     assignedGrade,
		reason:            reason,
		createdAt:         r.CreatedAt,
	}, nil
}
