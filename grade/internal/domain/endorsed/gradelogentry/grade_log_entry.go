package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewGradeLogEntry(
	endorsedId member.TenantMemberId,
	endorsedVersion uint,
	assignedGrade shared.Grade,
	reason Reason,
	createdAt time.Time,
) (GradeLogEntry, error) {
	return GradeLogEntry{
		endorsedId:      endorsedId,
		endorsedVersion: endorsedVersion,
		assignedGrade:   assignedGrade,
		reason:          reason,
		createdAt:       createdAt,
	}, nil
}

type GradeLogEntry struct {
	endorsedId      member.TenantMemberId
	endorsedVersion uint
	assignedGrade   shared.Grade
	reason          Reason
	createdAt       time.Time
}

func (gle GradeLogEntry) Export(ex GradeLogEntryExporterSetter) {
	ex.SetEndorsedId(gle.endorsedId)
	ex.SetEndorsedVersion(gle.endorsedVersion)
	ex.SetAssignedGrade(gle.assignedGrade)
	ex.SetReason(gle.reason)
	ex.SetCreatedAt(gle.createdAt)
}

type GradeLogEntryExporterSetter interface {
	SetEndorsedId(member.TenantMemberId)
	SetEndorsedVersion(uint)
	SetAssignedGrade(shared.Grade)
	SetReason(Reason)
	SetCreatedAt(time.Time)
}
