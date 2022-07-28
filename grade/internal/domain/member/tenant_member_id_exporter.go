package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewTenantMemberIdExporter(tenantId uint64, memberId uint64) *TenantMemberIdExporter {
	return &TenantMemberIdExporter{
		TenantId: seedwork.NewUint64Exporter(tenantId),
		MemberId: seedwork.NewUint64Exporter(memberId),
	}
}

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
