package competence

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func TestCompetenceExport(t *testing.T) {
	var actualExporter CompetenceExporter
	f := NewCompetenceFaker()
	agg, err := f.Create(nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, CompetenceExporter{
		Id:        values.NewCompetenceIdExporter(f.Id.TenantId, f.Id.CompetenceId),
		Name:      f.Name,
		OwnerId:   member.NewMemberIdExporter(f.OwnerId.TenantId, f.OwnerId.MemberId),
		CreatedAt: f.CreatedAt,
		Version:   1,
	}, actualExporter)
}
