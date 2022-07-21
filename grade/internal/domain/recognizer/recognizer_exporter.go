package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"time"
)

type RecognizerExporter struct {
	Id                        seedwork.Uint64Exporter
	UserId                    seedwork.Uint64Exporter
	Grade                     seedwork.Uint8Exporter
	AvailableEndorsementCount seedwork.Uint8Exporter
	Version                   uint
	CreatedAt                 time.Time
}

func (ex RecognizerExporter) SetState(
	id seedwork.Uint64Exporter,
	userId seedwork.Uint64Exporter,
	grade seedwork.Uint8Exporter,
	availableEndorsementCount seedwork.Uint8Exporter,
	version uint,
	createdAt time.Time,
) {
	ex.Id = id
	ex.UserId = userId
	ex.Grade = grade
	ex.AvailableEndorsementCount = availableEndorsementCount
	ex.Version = version
	ex.CreatedAt = createdAt
}
