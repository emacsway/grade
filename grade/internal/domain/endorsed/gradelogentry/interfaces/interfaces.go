package interfaces

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

type GradeLogEntryExporter interface {
	SetState(
		endorsedId interfaces.Exporter[uint64],
		endorsedVersion uint,
		assignedGrade interfaces.Exporter[uint8],
		reason interfaces.Exporter[string],
		createdAt time.Time,
	)
}
