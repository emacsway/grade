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
	id.ExportTo(&ex.Id)
}

func (ex *EndorsedExporter) AddEndorsement(ent endorsement.Endorsement) {
	var endorsementExporter endorsement.EndorsementExporter
	ent.ExportTo(&endorsementExporter)
	ex.ReceivedEndorsements = append(ex.ReceivedEndorsements, endorsementExporter)
}

func (ex *EndorsedExporter) AddGradeLogEntry(gle gradelogentry.GradeLogEntry) {
	var gradeLogEntryExporter gradelogentry.GradeLogEntryExporter
	gle.ExportTo(&gradeLogEntryExporter)
	ex.GradeLogEntries = append(ex.GradeLogEntries, gradeLogEntryExporter)
}

type EndorsedState struct {
	Id                   member.TenantMemberIdState
	Grade                uint8
	ReceivedEndorsements []endorsement.EndorsementState
	GradeLogEntries      []gradelogentry.GradeLogEntryState
	Version              uint
	CreatedAt            time.Time
}
