package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantArtifactId(tenantId, artifactId uint) (TenantArtifactId, error) {
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

func (cid TenantArtifactId) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(TenantArtifactId)
	if !ok {
		return false
	}
	return cid.tenantId.Equal(otherTyped.TenantId()) && cid.artifactId.Equal(otherTyped.ArtifactId())
}

func (cid TenantArtifactId) Export(ex TenantArtifactIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetArtifactId(cid.artifactId)
}

type TenantArtifactIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetArtifactId(ArtifactId)
}
