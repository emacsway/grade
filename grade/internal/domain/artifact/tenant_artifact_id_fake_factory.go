package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantArtifactIdFakeFactory() TenantArtifactIdFakeFactory {
	return TenantArtifactIdFakeFactory{
		TenantId:   tenant.TenantIdFakeValue,
		ArtifactId: uuid.ParseSilent("66e2fc13-89e3-483a-aa31-b8e75a20afba"),
	}
}

type TenantArtifactIdFakeFactory struct {
	TenantId   uuid.Uuid
	ArtifactId uuid.Uuid
}

func (f *TenantArtifactIdFakeFactory) NextTenantId() {
	f.TenantId = uuid.NewUuid()
}

func (f *TenantArtifactIdFakeFactory) NextArtifactId() {
	f.ArtifactId = uuid.NewUuid()
}

func (f TenantArtifactIdFakeFactory) Create() (TenantArtifactId, error) {
	return NewTenantArtifactId(f.TenantId, f.ArtifactId)
}
