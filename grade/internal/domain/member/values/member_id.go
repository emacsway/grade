package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewMemberId(tenantId, memberId uint) (MemberId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return MemberId{}, err
	}
	mId, err := NewInternalMemberId(memberId)
	if err != nil {
		return MemberId{}, err
	}
	return MemberId{
		tenantId: tId,
		memberId: mId,
	}, nil
}

func NewTransientMemberId(tenantId uint) (MemberId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return MemberId{}, err
	}
	mId := NewTransientInternalMemberId()
	return MemberId{
		tenantId: tId,
		memberId: mId,
	}, nil
}

type MemberId struct {
	tenantId tenant.TenantId
	memberId InternalMemberId
}

func (cid MemberId) TenantId() tenant.TenantId {
	return cid.tenantId
}

func (cid MemberId) MemberId() InternalMemberId {
	return cid.memberId
}

func (cid MemberId) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(MemberId)
	if !ok {
		return false
	}
	return cid.tenantId.Equal(otherTyped.TenantId()) && cid.memberId.Equal(otherTyped.MemberId())
}

func (cid MemberId) Export(ex TenantMemberIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetMemberId(cid.memberId)
}

type TenantMemberIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetMemberId(InternalMemberId)
}
