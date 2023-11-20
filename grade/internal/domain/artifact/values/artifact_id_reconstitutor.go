package values

type ArtifactIdReconstitutor struct {
	TenantId   uint
	ArtifactId uint
}

func (r ArtifactIdReconstitutor) Reconstitute() (ArtifactId, error) {
	return NewArtifactId(r.TenantId, r.ArtifactId)
}
