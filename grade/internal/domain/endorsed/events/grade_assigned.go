package events

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/assignment"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

func NewGradeAssigned(
	endorsedId member.TenantMemberId,
	endorsedVersion uint,
	assignedGrade grade.Grade,
	reason assignment.Reason,
	createdAt time.Time,
) GradeAssigned {
	return GradeAssigned{
		endorsedId:      endorsedId,
		endorsedVersion: endorsedVersion,
		assignedGrade:   assignedGrade,
		reason:          reason,
		createdAt:       createdAt,
	}
}

type GradeAssigned struct {
	endorsedId      member.TenantMemberId
	endorsedVersion uint
	assignedGrade   grade.Grade
	reason          assignment.Reason
	createdAt       time.Time
}

func (e GradeAssigned) EndorsedId() member.TenantMemberId {
	return e.endorsedId
}

func (e GradeAssigned) EndorsedVersion() uint {
	return e.endorsedVersion
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
