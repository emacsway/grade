package values

type MemberIdReconstitutor struct {
	TenantId uint
	MemberId uint
}

func (r MemberIdReconstitutor) Reconstitute() (MemberId, error) {
	return NewMemberId(r.TenantId, r.MemberId)
}
