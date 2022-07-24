package interfaces

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type RecognizerExporter interface {
	SetState(
		id interfaces.Exporter[uint64],
		memberId interfaces.Exporter[uint64],
		grade interfaces.Exporter[uint8],
		availableEndorsementCount interfaces.Exporter[uint8],
		pendingEndorsementCount interfaces.Exporter[uint8],
		version uint,
		createdAt time.Time,
	)
}
