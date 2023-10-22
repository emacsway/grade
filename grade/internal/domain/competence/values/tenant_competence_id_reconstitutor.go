package values

type TenantCompetenceIdReconstitutor struct {
	TenantId             uint
	CompetenceInTenantId uint
}

func (r TenantCompetenceIdReconstitutor) Reconstitute() (TenantCompetenceId, error) {
	return NewTenantCompetenceId(r.TenantId, r.CompetenceInTenantId)
}
