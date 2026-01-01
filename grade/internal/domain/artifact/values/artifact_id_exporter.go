package values

import (
	tenant "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func NewArtifactIdExporter(tenantId, artifactId uint) ArtifactIdExporter {
	return ArtifactIdExporter{
		TenantId:   exporters.UintExporter(tenantId),
		ArtifactId: exporters.UintExporter(artifactId),
	}
}

type ArtifactIdExporter struct {
	TenantId   exporters.UintExporter
	ArtifactId exporters.UintExporter
}

func (ex *ArtifactIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *ArtifactIdExporter) SetArtifactId(val InternalArtifactId) {
	val.Export(&ex.ArtifactId)
}
