package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

type EndorsedExporter struct {
	Id                   member.TenantMemberIdExporter
	Grade                seedwork.Uint8Exporter
	ReceivedEndorsements []endorsement.EndorsementExporter
	GradeLogEntries      []gradelogentry.GradeLogEntryExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetId(id member.TenantMemberId) {
	id.Export(&ex.Id)
}

func (ex *EndorsedExporter) SetGrade(g shared.Grade) {
	g.Export(&ex.Grade)
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

func (ex *EndorsedExporter) SetVersion(version uint) {
	ex.Version = version
}

func (ex *EndorsedExporter) SetCreatedAt(createdAt time.Time) {
	ex.CreatedAt = createdAt
}
