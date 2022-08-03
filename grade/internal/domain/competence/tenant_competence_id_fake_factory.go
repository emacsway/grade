package competence

func NewTenantCompetenceIdFakeFactory() TenantCompetenceIdFakeFactory {
	return TenantCompetenceIdFakeFactory{
		TenantId: 10,
		CompetenceId: 1,
	}
}

type TenantCompetenceIdFakeFactory struct {
	TenantId uint64
	CompetenceId uint64
}

func (f TenantCompetenceIdFakeFactory) Create() (TenantCompetenceId, error) {
	return NewTenantCompetenceId(f.TenantId, f.CompetenceId)
}
