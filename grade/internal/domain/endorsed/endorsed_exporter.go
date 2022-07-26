package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	interfaces3 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

type EndorsedExporter struct {
	Id                   interfaces.Exporter[uint64]
	Grade                interfaces.Exporter[uint8]
	ReceivedEndorsements []interfaces2.EndorsementExporter
	GradeLogEntries      []interfaces3.GradeLogEntryExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetState(
	id interfaces.Exporter[uint64],
	grade interfaces.Exporter[uint8],
	receivedEndorsements []interfaces2.EndorsementExporter,
	gradeLogEntries []interfaces3.GradeLogEntryExporter,
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
	Id                   uint64
	Grade                uint8
	ReceivedEndorsements []endorsement.EndorsementState
	GradeLogEntries      []gradelogentry.GradeLogEntryState
	Version              uint
	CreatedAt            time.Time
}
