package interfaces

import (
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsedExporter interface {
	SetState(
		id interfaces.Exporter[uint64],
		userId interfaces.Exporter[uint64],
		grade interfaces.Exporter[uint8],
		receivedEndorsements []interfaces2.EndorsementExporter,
		version uint,
		createdAt time.Time,
	)
}
