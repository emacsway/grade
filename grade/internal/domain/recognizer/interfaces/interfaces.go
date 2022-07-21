package interfaces

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"time"
)

type RecognizerExporter interface {
	SetState(
		id seedwork.Uint64Exporter,
		userId seedwork.Uint64Exporter,
		grade seedwork.Uint8Exporter,
		availableEndorsementCount seedwork.Uint8Exporter,
		version uint,
		createdAt time.Time,
	)
}
