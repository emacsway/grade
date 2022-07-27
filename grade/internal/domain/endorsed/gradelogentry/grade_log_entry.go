package gradelogentry

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"time"

	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewGradeLogEntry(
	endorsedId member.MemberId,
	endorsedVersion uint,
	assignedGrade shared.Grade,
	reason gradelogentry.Reason,
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
	endorsedId      member.MemberId
	endorsedVersion uint
	assignedGrade   shared.Grade
	reason          gradelogentry.Reason
	createdAt       time.Time
}

func (gle GradeLogEntry) ExportTo(ex interfaces2.GradeLogEntryExporter) {
	var endorsedId seedwork.Uint64Exporter
	var assignedGrade seedwork.Uint8Exporter
	var reason seedwork.StringExporter

	gle.endorsedId.ExportTo(&endorsedId)
	gle.assignedGrade.ExportTo(&assignedGrade)
	gle.reason.ExportTo(&reason)
	ex.SetState(
		&endorsedId, gle.endorsedVersion, &assignedGrade, &reason, gle.createdAt,
	)
}

func (gle GradeLogEntry) Export() GradeLogEntryState {
	return GradeLogEntryState{
		gle.endorsedId.Export(), gle.endorsedVersion,
		gle.assignedGrade.Export(), gle.reason.Export(), gle.createdAt,
	}
}
