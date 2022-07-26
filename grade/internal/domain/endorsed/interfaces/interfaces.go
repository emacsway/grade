package interfaces

import (
	"time"

	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	interfaces3 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

type EndorsedExporter interface {
	SetState(
		id interfaces.Exporter[uint64],
		grade interfaces.Exporter[uint8],
		receivedEndorsements []interfaces2.EndorsementExporter,
		gradeLogEntries []interfaces3.GradeLogEntryExporter,
		version uint,
		createdAt time.Time,
	)
}
