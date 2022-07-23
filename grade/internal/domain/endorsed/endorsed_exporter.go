package endorsed

import (
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsedExporter struct {
	Id                   interfaces.Exporter[uint64]
	MemberId             interfaces.Exporter[uint64]
	Grade                interfaces.Exporter[uint8]
	ReceivedEndorsements []interfaces2.EndorsementExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetState(
	id interfaces.Exporter[uint64],
	memberId interfaces.Exporter[uint64],
	grade interfaces.Exporter[uint8],
	receivedEndorsements []interfaces2.EndorsementExporter,
	version uint,
	createdAt time.Time,
) {
	ex.Id = id
	ex.MemberId = memberId
	ex.Grade = grade
	ex.ReceivedEndorsements = receivedEndorsements
	ex.Version = version
	ex.CreatedAt = createdAt
}
