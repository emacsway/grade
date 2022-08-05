package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantArtifactIdExporter(tenantId, artifactId uuid.Uuid) TenantArtifactIdExporter {
	return TenantArtifactIdExporter{
		TenantId:   exporters.UuidExporter(tenantId),
		ArtifactId: exporters.UuidExporter(artifactId),
	}
}

type TenantArtifactIdExporter struct {
	TenantId   exporters.UuidExporter
	ArtifactId exporters.UuidExporter
}

func (ex *TenantArtifactIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantArtifactIdExporter) SetArtifactId(val ArtifactId) {
	val.Export(&ex.ArtifactId)
}
