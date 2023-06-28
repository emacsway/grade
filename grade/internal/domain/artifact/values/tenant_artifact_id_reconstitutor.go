package values

type TenantArtifactIdReconstitutor struct {
	TenantId   uint
	ArtifactId uint
}

func (r TenantArtifactIdReconstitutor) Reconstitute() (TenantArtifactId, error) {
	return NewTenantArtifactId(r.TenantId, r.ArtifactId)
}
