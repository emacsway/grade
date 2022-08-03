package artifact

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/tenant"
)

func NewTenantArtifactIdExporter(tenantId, artifactId uint64) TenantArtifactIdExporter {
	return TenantArtifactIdExporter{
		TenantId:   seedwork.Uint64Exporter(tenantId),
		ArtifactId: seedwork.Uint64Exporter(artifactId),
	}
}

type TenantArtifactIdExporter struct {
	TenantId   seedwork.Uint64Exporter
	ArtifactId seedwork.Uint64Exporter
}

func (ex *TenantArtifactIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantArtifactIdExporter) SetArtifactId(val ArtifactId) {
	val.Export(&ex.ArtifactId)
}
