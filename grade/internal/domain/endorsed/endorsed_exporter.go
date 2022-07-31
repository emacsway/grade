package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type EndorsedExporter struct {
	Id                   member.TenantMemberIdExporter
	Grade                seedwork.Uint8Exporter
	ReceivedEndorsements []endorsement.EndorsementExporter
	GradeLogEntries      []gradelogentry.GradeLogEntryExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetId(val member.TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *EndorsedExporter) SetGrade(val grade.Grade) {
	val.Export(&ex.Grade)
}

func (ex *EndorsedExporter) AddEndorsement(val endorsement.Endorsement) {
	var endorsementExporter endorsement.EndorsementExporter
	val.Export(&endorsementExporter)
	ex.ReceivedEndorsements = append(ex.ReceivedEndorsements, endorsementExporter)
}

func (ex *EndorsedExporter) AddGradeLogEntry(val gradelogentry.GradeLogEntry) {
	var gradeLogEntryExporter gradelogentry.GradeLogEntryExporter
	val.Export(&gradeLogEntryExporter)
	ex.GradeLogEntries = append(ex.GradeLogEntries, gradeLogEntryExporter)
}

func (ex *EndorsedExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *EndorsedExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
