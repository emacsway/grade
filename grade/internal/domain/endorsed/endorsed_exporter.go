package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type EndorsedExporter struct {
	Id                   member.TenantMemberIdExporter
	Grade                seedwork.ExporterSetter[uint8]
	ReceivedEndorsements []endorsement.EndorsementExporter
	GradeLogEntries      []gradelogentry.GradeLogEntryExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetState(
	grade seedwork.ExporterSetter[uint8],
	version uint,
	createdAt time.Time,
) {
	ex.Grade = grade
	ex.Version = version
	ex.CreatedAt = createdAt
}

func (ex *EndorsedExporter) SetId(id member.TenantMemberId) {
	id.Export(&ex.Id)
}

func (ex *EndorsedExporter) AddEndorsement(ent endorsement.Endorsement) {
	var endorsementExporter endorsement.EndorsementExporter
	ent.Export(&endorsementExporter)
	ex.ReceivedEndorsements = append(ex.ReceivedEndorsements, endorsementExporter)
}

func (ex *EndorsedExporter) AddGradeLogEntry(gle gradelogentry.GradeLogEntry) {
	var gradeLogEntryExporter gradelogentry.GradeLogEntryExporter
	gle.Export(&gradeLogEntryExporter)
	ex.GradeLogEntries = append(ex.GradeLogEntries, gradeLogEntryExporter)
}
