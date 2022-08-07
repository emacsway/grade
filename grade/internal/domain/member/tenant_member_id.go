package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantMemberId(tenantId, memberId uuid.Uuid) (TenantMemberId, error) {
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

func (cid TenantMemberId) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(TenantMemberId)
	if !ok {
		return false
	}
	return cid.tenantId.Equal(otherTyped.TenantId()) && cid.memberId.Equal(otherTyped.MemberId())
}

func (cid TenantMemberId) Export(ex TenantMemberIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetMemberId(cid.memberId)
}

type TenantMemberIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetMemberId(MemberId)
}
