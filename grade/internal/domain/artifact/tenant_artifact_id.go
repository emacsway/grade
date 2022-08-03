package artifact

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/tenant"
)

func NewTenantArtifactId(tenantId, artifactId uint64) (TenantArtifactId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return TenantArtifactId{}, err
	}
	mId, err := NewArtifactId(artifactId)
	if err != nil {
		return TenantArtifactId{}, err
	}
	return TenantArtifactId{
		tenantId:   tId,
		artifactId: mId,
	}, nil
}

type TenantArtifactId struct {
	tenantId   tenant.TenantId
	artifactId ArtifactId
}

func (cid TenantArtifactId) TenantId() tenant.TenantId {
	return cid.tenantId
}

func (cid TenantArtifactId) ArtifactId() ArtifactId {
	return cid.artifactId
}

func (cid TenantArtifactId) Equal(other TenantArtifactId) bool {
	return cid.tenantId.Equal(other.TenantId()) && cid.artifactId.Equal(other.ArtifactId())
}

func (cid TenantArtifactId) Export(ex TenantArtifactIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetArtifactId(cid.artifactId)
}

type TenantArtifactIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetArtifactId(ArtifactId)
}
