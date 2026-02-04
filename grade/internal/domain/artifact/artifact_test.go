package artifact

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func TestArtifactExport(t *testing.T) {
	var actualExporter ArtifactExporter
	f := NewArtifactFaker()
	agg, err := f.Create(nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	var expectedStatus uint8
	f.Status.Export(func(v uint8) { expectedStatus = v })
	assert.Equal(t, ArtifactExporter{
		Id:          values.NewArtifactIdExporter(f.Id.TenantId, f.Id.ArtifactId),
		Status:      expectedStatus,
		Name:        f.Name,
		Description: f.Description,
		Url:         f.Url,
		CompetenceIds: []competence.CompetenceIdExporter{
			competence.NewCompetenceIdExporter(
				f.CompetenceIds[0].TenantId,
				f.CompetenceIds[0].CompetenceId,
			),
		},
		OwnerId:   member.NewMemberIdExporter(f.OwnerId.TenantId, f.OwnerId.MemberId),
		CreatedAt: f.CreatedAt,
		Version:   1,
	}, actualExporter)
}
