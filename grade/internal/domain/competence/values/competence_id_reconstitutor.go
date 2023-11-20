package values

type CompetenceIdReconstitutor struct {
	TenantId     uint
	CompetenceId uint
}

func (r CompetenceIdReconstitutor) Reconstitute() (CompetenceId, error) {
	return NewCompetenceId(r.TenantId, r.CompetenceId)
}
