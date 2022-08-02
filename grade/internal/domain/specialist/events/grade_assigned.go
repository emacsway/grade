package events

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/assignment"
)

func NewGradeAssigned(
	specialistId member.TenantMemberId,
	specialistVersion uint,
	assignedGrade grade.Grade,
	reason assignment.Reason,
	createdAt time.Time,
) GradeAssigned {
	return GradeAssigned{
		specialistId:      specialistId,
		specialistVersion: specialistVersion,
		assignedGrade:     assignedGrade,
		reason:            reason,
		createdAt:         createdAt,
	}
}

type GradeAssigned struct {
	specialistId      member.TenantMemberId
	specialistVersion uint
	assignedGrade     grade.Grade
	reason            assignment.Reason
	createdAt         time.Time
}

func (e GradeAssigned) SpecialistId() member.TenantMemberId {
	return e.specialistId
}

func (e GradeAssigned) SpecialistVersion() uint {
	return e.specialistVersion
}

func (e GradeAssigned) AssignedGrade() grade.Grade {
	return e.assignedGrade
}

func (e GradeAssigned) Reason() assignment.Reason {
	return e.reason
}

func (e GradeAssigned) CreatedAt() time.Time {
	return e.createdAt
}
