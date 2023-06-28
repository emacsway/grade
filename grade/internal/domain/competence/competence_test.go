package competence

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestCompetenceExport(t *testing.T) {
	var actualExporter CompetenceExporter
	f := NewCompetenceFaker()
	agg, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, CompetenceExporter{
		Id:        values.NewTenantCompetenceIdExporter(f.Id.TenantId, f.Id.CompetenceId),
		Name:      exporters.StringExporter(f.Name),
		OwnerId:   member.NewTenantMemberIdExporter(f.OwnerId.TenantId, f.OwnerId.MemberId),
		CreatedAt: f.CreatedAt,
	}, actualExporter)
}
