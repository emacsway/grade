package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
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

func (gle GradeLogEntry) ExportTo(ex GradeLogEntryExporterSetter) {
	var assignedGrade seedwork.Uint8Exporter
	var reason seedwork.StringExporter

	gle.assignedGrade.ExportTo(&assignedGrade)
	gle.reason.ExportTo(&reason)
	ex.SetState(
		gle.endorsedVersion, &assignedGrade, &reason, gle.createdAt,
	)
	ex.SetEndorsedId(gle.endorsedId)
}

func (gle GradeLogEntry) Export() GradeLogEntryState {
	return GradeLogEntryState{
		gle.endorsedId.Export(), gle.endorsedVersion,
		gle.assignedGrade.Export(), gle.reason.Export(), gle.createdAt,
	}
}

type GradeLogEntryExporterSetter interface {
	SetState(
		endorsedVersion uint,
		assignedGrade seedwork.ExporterSetter[uint8],
		reason seedwork.ExporterSetter[string],
		createdAt time.Time,
	)
	SetEndorsedId(member.TenantMemberId)
}
