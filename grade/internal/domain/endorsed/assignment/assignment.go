package assignment

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

func NewAssignment(
	endorsedId member.TenantMemberId,
	endorsedVersion uint,
	assignedGrade grade.Grade,
	reason Reason,
	createdAt time.Time,
) (Assignment, error) {
	return Assignment{
		endorsedId:      endorsedId,
		endorsedVersion: endorsedVersion,
		assignedGrade:   assignedGrade,
		reason:          reason,
		createdAt:       createdAt,
	}, nil
}

type Assignment struct {
	endorsedId      member.TenantMemberId
	endorsedVersion uint
	assignedGrade   grade.Grade
	reason          Reason
	createdAt       time.Time
}

func (a Assignment) Export(ex AssignmentExporterSetter) {
	ex.SetEndorsedId(a.endorsedId)
	ex.SetEndorsedVersion(a.endorsedVersion)
	ex.SetAssignedGrade(a.assignedGrade)
	ex.SetReason(a.reason)
	ex.SetCreatedAt(a.createdAt)
}

type AssignmentExporterSetter interface {
	SetEndorsedId(member.TenantMemberId)
	SetEndorsedVersion(uint)
	SetAssignedGrade(grade.Grade)
	SetReason(Reason)
	SetCreatedAt(time.Time)
}
