package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantArtifactIdFakeFactory() TenantArtifactIdFakeFactory {
	return TenantArtifactIdFakeFactory{
		TenantId:   tenant.TenantIdFakeValue,
		ArtifactId: uint(3),
	}
}

type TenantArtifactIdFakeFactory struct {
	TenantId   uint
	ArtifactId uint
}

func (f *TenantArtifactIdFakeFactory) NextTenantId() {
	f.TenantId += 1
}

func (f *TenantArtifactIdFakeFactory) NextArtifactId() {
	f.ArtifactId += 1
}

func (f TenantArtifactIdFakeFactory) Create() (TenantArtifactId, error) {
	return NewTenantArtifactId(f.TenantId, f.ArtifactId)
}
