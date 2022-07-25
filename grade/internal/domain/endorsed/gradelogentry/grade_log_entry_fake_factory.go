package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewGradeLogEntryFakeFactory() *GradeLogEntryFakeFactory {
	return &GradeLogEntryFakeFactory{
		1, 2, 1, time.Now(),
	}
}

type GradeLogEntryFakeFactory struct {
	EndorsedId      uint64
	EndorsedVersion uint
	AssignedGrade   uint8
	CreatedAt       time.Time
}

func (f GradeLogEntryFakeFactory) Create() (GradeLogEntry, error) {
	endorsedId, _ := endorsed.NewEndorsedId(f.EndorsedId)
	assignedGrade, _ := shared.NewGrade(f.AssignedGrade)
	return NewGradeLogEntry(endorsedId, f.EndorsedVersion, assignedGrade, f.CreatedAt)
}

func (f GradeLogEntryFakeFactory) Export() GradeLogEntryState {
	return GradeLogEntryState{
		f.EndorsedId, f.EndorsedVersion, f.AssignedGrade, f.CreatedAt,
	}
}

func (f GradeLogEntryFakeFactory) ExportTo(ex interfaces.GradeLogEntryExporter) {
	var endorsedId seedwork.Uint64Exporter
	var assignedGrade seedwork.Uint8Exporter

	endorsedId.SetState(f.EndorsedId)
	assignedGrade.SetState(f.AssignedGrade)
	ex.SetState(
		&endorsedId, f.EndorsedVersion, &assignedGrade, f.CreatedAt,
	)
}
