package gradelogentry

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func TestRecognizerExport(t *testing.T) {
	var actualExporter GradeLogEntryExporter
	f := NewGradeLogEntryFakeFactory()
	agg, _ := f.Create()
	agg.Export(&actualExporter)
	assert.Equal(t, GradeLogEntryExporter{
		EndorsedId:      member.NewTenantMemberIdExporter(f.EndorsedId.TenantId, f.EndorsedId.MemberId),
		EndorsedVersion: f.EndorsedVersion,
		AssignedGrade:   seedwork.Uint8Exporter(f.AssignedGrade),
		Reason:          seedwork.StringExporter(f.Reason),
		CreatedAt:       f.CreatedAt,
	}, actualExporter)
}
