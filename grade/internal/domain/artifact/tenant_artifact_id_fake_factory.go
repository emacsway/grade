package artifact

func NewTenantArtifactIdFakeFactory() TenantArtifactIdFakeFactory {
	return TenantArtifactIdFakeFactory{
		TenantId:   10,
		ArtifactId: 1,
	}
}

type TenantArtifactIdFakeFactory struct {
	TenantId   uint64
	ArtifactId uint64
}

func (f TenantArtifactIdFakeFactory) Create() (TenantArtifactId, error) {
	return NewTenantArtifactId(f.TenantId, f.ArtifactId)
}
