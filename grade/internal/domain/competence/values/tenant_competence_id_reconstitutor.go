package values

type TenantCompetenceIdReconstitutor struct {
	TenantId     uint
	CompetenceId uint
}

func (r TenantCompetenceIdReconstitutor) Reconstitute() (TenantCompetenceId, error) {
	return NewTenantCompetenceId(r.TenantId, r.CompetenceId)
}
