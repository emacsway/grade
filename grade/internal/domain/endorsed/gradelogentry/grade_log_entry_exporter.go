package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type GradeLogEntryExporter struct {
	EndorsedId      member.TenantMemberIdExporter
	EndorsedVersion uint
	AssignedGrade   seedwork.Uint8Exporter
	Reason          seedwork.StringExporter
	CreatedAt       time.Time
}

func (ex *GradeLogEntryExporter) SetEndorsedId(val member.TenantMemberId) {
	val.Export(&ex.EndorsedId)
}

func (ex *GradeLogEntryExporter) SetEndorsedVersion(val uint) {
	ex.EndorsedVersion = val
}

func (ex *GradeLogEntryExporter) SetAssignedGrade(val grade.Grade) {
	val.Export(&ex.AssignedGrade)
}

func (ex *GradeLogEntryExporter) SetReason(val Reason) {
	val.Export(&ex.Reason)
}

func (ex *GradeLogEntryExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
