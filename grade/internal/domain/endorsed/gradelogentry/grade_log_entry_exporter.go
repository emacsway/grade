package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type GradeLogEntryExporter struct {
	EndorsedId      member.TenantMemberIdExporter
	EndorsedVersion uint
	AssignedGrade   seedwork.ExporterSetter[uint8]
	Reason          seedwork.ExporterSetter[string]
	CreatedAt       time.Time
}

func (ex *GradeLogEntryExporter) SetState(
	endorsedVersion uint,
	assignedGrade seedwork.ExporterSetter[uint8],
	reason seedwork.ExporterSetter[string],
	createdAt time.Time,
) {
	ex.EndorsedVersion = endorsedVersion
	ex.AssignedGrade = assignedGrade
	ex.Reason = reason
	ex.CreatedAt = createdAt
}

func (ex *GradeLogEntryExporter) SetEndorsedId(endorsedId member.TenantMemberId) {
	endorsedId.ExportTo(&ex.EndorsedId)
}
