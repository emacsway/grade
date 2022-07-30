package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type EndorsedExporter struct {
	Id                   member.TenantMemberIdExporterSetter
	Grade                seedwork.ExporterSetter[uint8]
	ReceivedEndorsements []endorsement.EndorsementExporterSetter
	GradeLogEntries      []gradelogentry.GradeLogEntryExporterSetter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetState(
	id member.TenantMemberIdExporterSetter,
	grade seedwork.ExporterSetter[uint8],
	receivedEndorsements []endorsement.EndorsementExporterSetter,
	gradeLogEntries []gradelogentry.GradeLogEntryExporterSetter,
	version uint,
	createdAt time.Time,
) {
	ex.Id = id
	ex.Grade = grade
	ex.ReceivedEndorsements = receivedEndorsements
	ex.GradeLogEntries = gradeLogEntries
	ex.Version = version
	ex.CreatedAt = createdAt
}

type EndorsedState struct {
	Id                   member.TenantMemberIdState
	Grade                uint8
	ReceivedEndorsements []endorsement.EndorsementState
	GradeLogEntries      []gradelogentry.GradeLogEntryState
	Version              uint
	CreatedAt            time.Time
}
