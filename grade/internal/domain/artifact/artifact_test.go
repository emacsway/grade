package artifact

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestArtifactExport(t *testing.T) {
	var actualExporter ArtifactExporter
	f := NewArtifactFakeFactory()
	agg, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, ArtifactExporter{
		Id:          NewTenantArtifactIdExporter(f.Id.TenantId, f.Id.ArtifactId),
		Status:      exporters.Uint8Exporter(f.Status),
		Name:        exporters.StringExporter(f.Name),
		Description: exporters.StringExporter(f.Description),
		Url:         exporters.StringExporter(f.Url),
		CompetenceIds: []competence.TenantCompetenceIdExporter{
			competence.NewTenantCompetenceIdExporter(
				f.CompetenceIds[0].TenantId,
				f.CompetenceIds[0].CompetenceId,
			),
		},
		OwnerId:   member.NewTenantMemberIdExporter(f.OwnerId.TenantId, f.OwnerId.MemberId),
		CreatedAt: f.CreatedAt,
	}, actualExporter)
}
