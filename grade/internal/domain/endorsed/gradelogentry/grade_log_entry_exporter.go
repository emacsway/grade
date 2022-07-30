package gradelogentry

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"time"
)

type GradeLogEntryExporter struct {
	EndorsedId      member.TenantMemberIdExporterSetter
	EndorsedVersion uint
	AssignedGrade   seedwork.ExporterSetter[uint8]
	Reason          seedwork.ExporterSetter[string]
	CreatedAt       time.Time
}

func (ex *GradeLogEntryExporter) SetState(
	endorsedId member.TenantMemberIdExporterSetter,
	endorsedVersion uint,
	assignedGrade seedwork.ExporterSetter[uint8],
	reason seedwork.ExporterSetter[string],
	createdAt time.Time,
) {
	ex.EndorsedId = endorsedId
	ex.EndorsedVersion = endorsedVersion
	ex.AssignedGrade = assignedGrade
	ex.Reason = reason
	ex.CreatedAt = createdAt
}

type GradeLogEntryState struct {
	EndorsedId      member.TenantMemberIdState
	EndorsedVersion uint
	AssignedGrade   uint8
	Reason          string
	CreatedAt       time.Time
}
