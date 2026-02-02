package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/specification/domain"
)

func NewArtifactId(tenantId, artifactId uint) (ArtifactId, error) {
	tId, err := tenant.NewTenantId(tenantId)
	if err != nil {
		return ArtifactId{}, err
	}
	mId, err := NewInternalArtifactId(artifactId)
	if err != nil {
		return ArtifactId{}, err
	}
	return ArtifactId{
		tenantId:   tId,
		artifactId: mId,
	}, nil
}

type ArtifactId struct {
	tenantId   tenant.TenantId
	artifactId InternalArtifactId
}

func (cid ArtifactId) TenantId() tenant.TenantId {
	return cid.tenantId
}

func (cid ArtifactId) ArtifactId() InternalArtifactId {
	return cid.artifactId
}

func (cid ArtifactId) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(ArtifactId)
	if !ok {
		return false
	}
	return cid.tenantId.Equal(otherTyped.TenantId()) && cid.artifactId.Equal(otherTyped.ArtifactId())
}

func (cid ArtifactId) Export(ex ArtifactIdExporterSetter) {
	ex.SetTenantId(cid.tenantId)
	ex.SetArtifactId(cid.artifactId)
}

type ArtifactIdExporterSetter interface {
	SetTenantId(tenant.TenantId)
	SetArtifactId(InternalArtifactId)
}
