package assignment

import (
	"testing"

	"github.com/stretchr/testify/assert"

	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func TestAssignmentExport(t *testing.T) {
	var actualExporter AssignmentExporter
	f := NewAssignmentFaker()
	agg, _ := f.Create()
	agg.Export(&actualExporter)
	assert.Equal(t, AssignmentExporter{
		SpecialistId:      member.NewMemberIdExporter(f.SpecialistId.TenantId, f.SpecialistId.MemberId),
		SpecialistVersion: f.SpecialistVersion,
		AssignedGrade:     exporters.Uint8Exporter(f.AssignedGrade),
		Reason:            exporters.StringExporter(f.Reason),
		CreatedAt:         f.CreatedAt,
	}, actualExporter)
}
