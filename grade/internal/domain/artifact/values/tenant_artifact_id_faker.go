package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewArtifactIdFaker() ArtifactIdFaker {
	return ArtifactIdFaker{
		TenantId:   tenant.TenantIdFakeValue,
		ArtifactId: uint(3),
	}
}

type ArtifactIdFaker struct {
	TenantId   uint
	ArtifactId uint
}

func (f ArtifactIdFaker) Create() (ArtifactId, error) {
	return NewArtifactId(f.TenantId, f.ArtifactId)
}
