package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

type GradeLogEntryExporter struct {
	EndorsedId      interfaces.Exporter[uint64]
	EndorsedVersion uint
	AssignedGrade   interfaces.Exporter[uint8]
	CreatedAt       time.Time
}

func (ex *GradeLogEntryExporter) SetState(
	endorsedId interfaces.Exporter[uint64],
	endorsedVersion uint,
	assignedGrade interfaces.Exporter[uint8],
	createdAt time.Time,
) {
	ex.EndorsedId = endorsedId
	ex.EndorsedVersion = endorsedVersion
	ex.AssignedGrade = assignedGrade
	ex.CreatedAt = createdAt
}

type GradeLogEntryState struct {
	EndorsedId      uint64
	EndorsedVersion uint
	AssignedGrade   uint8
	CreatedAt       time.Time
}
