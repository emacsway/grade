package member

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/tenant"
)

func NewTenantMemberId(tenantId, memberId uint64) (TenantMemberId, error) {
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

func (cid TenantMemberId) Equal(other TenantMemberId) bool {
	return cid.tenantId.Equal(other.TenantId()) && cid.memberId.Equal(other.MemberId())
}

func (cid TenantMemberId) Export(ex TenantMemberIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetMemberId(cid.memberId)
}

type TenantMemberIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetMemberId(MemberId)
}
