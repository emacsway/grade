package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
)

func NewTenantArtifactIdExporter(tenantId, artifactId uuid.Uuid) TenantArtifactIdExporter {
	return TenantArtifactIdExporter{
		TenantId:   seedwork.UuidExporter(tenantId),
		ArtifactId: seedwork.UuidExporter(artifactId),
	}
}

type TenantArtifactIdExporter struct {
	TenantId   seedwork.UuidExporter
	ArtifactId seedwork.UuidExporter
}

func (ex *TenantArtifactIdExporter) SetTenantId(val tenant.TenantId) {
	val.Export(&ex.TenantId)
}

func (ex *TenantArtifactIdExporter) SetArtifactId(val ArtifactId) {
	val.Export(&ex.ArtifactId)
}
