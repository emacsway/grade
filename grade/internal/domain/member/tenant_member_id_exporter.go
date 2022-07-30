package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewTenantMemberIdExporter(tenantId uint64, memberId uint64) *TenantMemberIdExporter {
	return &TenantMemberIdExporter{
		TenantId: seedwork.NewUint64Exporter(tenantId),
		MemberId: seedwork.NewUint64Exporter(memberId),
	}
}

type TenantMemberIdExporter struct {
	TenantId seedwork.ExporterSetter[uint64]
	MemberId seedwork.ExporterSetter[uint64]
}

func (ex *TenantMemberIdExporter) SetState(
	tenantId seedwork.ExporterSetter[uint64],
	memberId seedwork.ExporterSetter[uint64],
) {
	ex.TenantId = tenantId
	ex.MemberId = memberId
}

type TenantMemberIdState struct {
	TenantId uint64
	MemberId uint64
}
