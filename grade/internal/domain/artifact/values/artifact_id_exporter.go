package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
)

func NewArtifactIdExporter(tenantId, artifactId uint) ArtifactIdExporter {
	return ArtifactIdExporter{
		TenantId:   tenantId,
		ArtifactId: artifactId,
	}
}

type ArtifactIdExporter struct {
	TenantId   uint
	ArtifactId uint
}

func (ex *ArtifactIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(func(v uint) { ex.TenantId = v })
}

func (ex *ArtifactIdExporter) SetArtifactId(val InternalArtifactId) {
	val.Export(func(v uint) { ex.ArtifactId = v })
}
