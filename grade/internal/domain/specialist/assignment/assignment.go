package assignment

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment/values"
)

func NewAssignment(
	specialistId member.MemberId,
	specialistVersion uint,
	assignedGrade grade.Grade,
	reason values.Reason,
	createdAt time.Time,
) (Assignment, error) {
	return Assignment{
		specialistId:      specialistId,
		specialistVersion: specialistVersion,
		assignedGrade:     assignedGrade,
		reason:            reason,
		createdAt:         createdAt,
	}, nil
}

type Assignment struct {
	specialistId      member.MemberId
	specialistVersion uint
	assignedGrade     grade.Grade
	reason            values.Reason
	createdAt         time.Time
}

func (a Assignment) Export(ex AssignmentExporterSetter) {
	ex.SetSpecialistId(a.specialistId)
	ex.SetSpecialistVersion(a.specialistVersion)
	ex.SetAssignedGrade(a.assignedGrade)
	ex.SetReason(a.reason)
	ex.SetCreatedAt(a.createdAt)
}

type AssignmentExporterSetter interface {
	SetSpecialistId(member.MemberId)
	SetSpecialistVersion(uint)
	SetAssignedGrade(grade.Grade)
	SetReason(values.Reason)
	SetCreatedAt(time.Time)
}
