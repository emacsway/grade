package interfaces

import (
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	interfaces3 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	interfaces4 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsedExporter interface {
	SetState(
		id interfaces4.TenantMemberIdExporter,
		grade interfaces.Exporter[uint8],
		receivedEndorsements []interfaces2.EndorsementExporter,
		gradeLogEntries []interfaces3.GradeLogEntryExporter,
		version uint,
		createdAt time.Time,
	)
}
