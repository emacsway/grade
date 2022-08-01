package member

func NewTenantMemberIdFakeFactory() TenantMemberIdFakeFactory {
	return TenantMemberIdFakeFactory{
		TenantId: 10,
		MemberId: 1,
	}
}

type TenantMemberIdFakeFactory struct {
	TenantId uint64
	MemberId uint64
}

func (f TenantMemberIdFakeFactory) Create() (TenantMemberId, error) {
	return NewTenantMemberId(f.TenantId, f.MemberId)
}
