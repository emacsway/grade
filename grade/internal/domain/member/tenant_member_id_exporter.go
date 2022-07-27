package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

type TenantMemberIdExporter struct {
	TenantId interfaces.Exporter[uint64]
	MemberId interfaces.Exporter[uint64]
}

func (ex *TenantMemberIdExporter) SetState(
	tenantId interfaces.Exporter[uint64],
	memberId interfaces.Exporter[uint64],
) {
	ex.TenantId = tenantId
	ex.MemberId = memberId
}

type TenantMemberIdState struct {
	TenantId uint64
	MemberId uint64
}
