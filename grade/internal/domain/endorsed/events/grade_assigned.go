package events

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/assignment"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

type GradeAssigned struct {
	EndorsedId      member.TenantMemberId
	EndorsedVersion uint
	AssignedGrade   grade.Grade
	Reason          assignment.Reason
	CreatedAt       time.Time
}
