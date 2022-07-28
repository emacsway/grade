package interfaces

import (
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type RecognizerExporter interface {
	SetState(
		id interfaces2.TenantMemberIdExporter,
		grade interfaces.Exporter[uint8],
		availableEndorsementCount interfaces.Exporter[uint8],
		pendingEndorsementCount interfaces.Exporter[uint8],
		version uint,
		createdAt time.Time,
	)
}
