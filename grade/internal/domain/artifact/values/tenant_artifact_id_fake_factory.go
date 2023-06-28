package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewTenantArtifactIdFaker() TenantArtifactIdFaker {
	return TenantArtifactIdFaker{
		TenantId:   tenant.TenantIdFakeValue,
		ArtifactId: uint(3),
	}
}

type TenantArtifactIdFaker struct {
	TenantId   uint
	ArtifactId uint
}

func (f TenantArtifactIdFaker) Create() (TenantArtifactId, error) {
	return NewTenantArtifactId(f.TenantId, f.ArtifactId)
}
