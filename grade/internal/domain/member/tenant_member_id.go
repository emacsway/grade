package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/tenant"
)

func NewTenantMemberId(tenantId uint64, memberId uint64) (TenantMemberId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return TenantMemberId{}, err
	}
	mId, err := NewMemberId(memberId)
	if err != nil {
		return TenantMemberId{}, err
	}
	return TenantMemberId{
		tenantId: tId,
		memberId: mId,
	}, nil
}

type TenantMemberId struct {
	tenantId tenant.TenantId
	memberId MemberId
}

func (cid TenantMemberId) TenantId() tenant.TenantId {
	return cid.tenantId
}

func (cid TenantMemberId) MemberId() MemberId {
	return cid.memberId
}

func (cid TenantMemberId) Equals(other TenantMemberId) bool {
	return cid.tenantId.Equals(other.TenantId()) && cid.memberId.Equals(other.MemberId())
}

func (cid TenantMemberId) Export(ex TenantMemberIdExporterSetter) {
	var tenantId, memberId seedwork.Uint64Exporter

	cid.tenantId.Export(&tenantId)
	cid.memberId.Export(&memberId)
	ex.SetState(&tenantId, &memberId)
}

type TenantMemberIdExporterSetter interface {
	SetState(
		tenantId seedwork.ExporterSetter[uint64],
		memberId seedwork.ExporterSetter[uint64],
	)
}
