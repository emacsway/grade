package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type RecognizerExporter struct {
	Id                        interfaces.Exporter[uint64]
	Grade                     interfaces.Exporter[uint8]
	AvailableEndorsementCount interfaces.Exporter[uint8]
	PendingEndorsementCount   interfaces.Exporter[uint8]
	Version                   uint
	CreatedAt                 time.Time
}

func (ex *RecognizerExporter) SetState(
	id interfaces.Exporter[uint64],
	grade interfaces.Exporter[uint8],
	availableEndorsementCount interfaces.Exporter[uint8],
	pendingEndorsementCount interfaces.Exporter[uint8],
	version uint,
	createdAt time.Time,
) {
	ex.Id = id
	ex.Grade = grade
	ex.AvailableEndorsementCount = availableEndorsementCount
	ex.PendingEndorsementCount = pendingEndorsementCount
	ex.Version = version
	ex.CreatedAt = createdAt
}

type RecognizerState struct {
	Id                        uint64
	Grade                     uint8
	AvailableEndorsementCount uint8
	PendingEndorsementCount   uint8
	Version                   uint
	CreatedAt                 time.Time
}
