package gradelogentry

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

type GradeLogEntryExporter struct {
	EndorsedId      interfaces.Exporter[uint64]
	EndorsedVersion uint
	AssignedGrade   interfaces.Exporter[uint8]
	Reason          interfaces.Exporter[string]
	CreatedAt       time.Time
}

func (ex *GradeLogEntryExporter) SetState(
	endorsedId interfaces.Exporter[uint64],
	endorsedVersion uint,
	assignedGrade interfaces.Exporter[uint8],
	reason interfaces.Exporter[string],
	createdAt time.Time,
) {
	ex.EndorsedId = endorsedId
	ex.EndorsedVersion = endorsedVersion
	ex.AssignedGrade = assignedGrade
	ex.Reason = reason
	ex.CreatedAt = createdAt
}

type GradeLogEntryState struct {
	EndorsedId      uint64
	EndorsedVersion uint
	AssignedGrade   uint8
	Reason          string
	CreatedAt       time.Time
}
