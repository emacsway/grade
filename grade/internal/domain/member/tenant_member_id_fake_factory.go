package member

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewTenantMemberIdFakeFactory() TenantMemberIdFakeFactory {
	return TenantMemberIdFakeFactory{
		TenantId: uuid.ParseSilent("63e8d541-af30-4593-a8ac-761dc268926d"),
		MemberId: uuid.ParseSilent("cf9462cf-51d3-4c0a-b5ef-53b3dfccc7f6"),
	}
}

type TenantMemberIdFakeFactory struct {
	TenantId uuid.Uuid
	MemberId uuid.Uuid
}

func (f TenantMemberIdFakeFactory) Create() (TenantMemberId, error) {
	return NewTenantMemberId(f.TenantId, f.MemberId)
}
