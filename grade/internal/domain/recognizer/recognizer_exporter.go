package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type RecognizerExporter struct {
	Id                        interfaces.Exporter[uint64]
	UserId                    interfaces.Exporter[uint64]
	Grade                     interfaces.Exporter[uint8]
	AvailableEndorsementCount interfaces.Exporter[uint8]
	Version                   uint
	CreatedAt                 time.Time
}

func (ex *RecognizerExporter) SetState(
	id interfaces.Exporter[uint64],
	userId interfaces.Exporter[uint64],
	grade interfaces.Exporter[uint8],
	availableEndorsementCount interfaces.Exporter[uint8],
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
