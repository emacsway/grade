package interfaces

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

type TenantMemberIdExporter interface {
	SetState(
		tenantId interfaces.Exporter[uint64],
		memberId interfaces.Exporter[uint64],
	)
}
