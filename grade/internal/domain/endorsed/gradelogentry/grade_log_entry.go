package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewGradeLogEntry(
	endorsedId endorsed.EndorsedId,
	endorsedVersion uint,
	assignedGrade shared.Grade,
	createdAt time.Time,
) (GradeLogEntry, error) {
	return GradeLogEntry{
		endorsedId, endorsedVersion, assignedGrade, createdAt,
	}, nil
}

type GradeLogEntry struct {
	endorsedId      endorsed.EndorsedId
	endorsedVersion uint
	assignedGrade   shared.Grade
	createdAt       time.Time
}

func (gle GradeLogEntry) ExportTo(ex interfaces2.GradeLogEntryExporter) {
	var endorsedId seedwork.Uint64Exporter
	var assignedGrade seedwork.Uint8Exporter

	gle.endorsedId.ExportTo(&endorsedId)
	gle.assignedGrade.ExportTo(&assignedGrade)
	ex.SetState(
		&endorsedId, gle.endorsedVersion, &assignedGrade, gle.createdAt,
	)
}

func (gle GradeLogEntry) Export() GradeLogEntryState {
	return GradeLogEntryState{
		gle.endorsedId.Export(), gle.endorsedVersion,
		gle.assignedGrade.Export(), gle.createdAt,
	}
}