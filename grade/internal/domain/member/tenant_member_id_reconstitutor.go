package member

type TenantMemberIdReconstitutor struct {
	TenantId uint
	MemberId uint
}

func (r TenantMemberIdReconstitutor) Reconstitute() (TenantMemberId, error) {
	return NewTenantMemberId(r.TenantId, r.MemberId)
}
