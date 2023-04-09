package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantArtifactIdExporter(tenantId, artifactId uint) TenantArtifactIdExporter {
	return TenantArtifactIdExporter{
		TenantId:   exporters.UintExporter(tenantId),
		ArtifactId: exporters.UintExporter(artifactId),
	}
}

type TenantArtifactIdExporter struct {
	TenantId   exporters.UintExporter
	ArtifactId exporters.UintExporter
}

func (ex *TenantArtifactIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantArtifactIdExporter) SetArtifactId(val ArtifactId) {
	val.Export(&ex.ArtifactId)
}
