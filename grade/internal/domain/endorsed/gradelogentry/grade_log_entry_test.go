package gradelogentry

import (
	"testing"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
)

func TestGradeLogEntryExport(t *testing.T) {
	f, err := NewGradeLogEntryFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	e, _ := f.Create()
	assert.Equal(t, GradeLogEntryState{
		EndorsedId:      f.EndorsedId,
		EndorsedVersion: f.EndorsedVersion,
		AssignedGrade:   f.AssignedGrade,
		CreatedAt:       f.CreatedAt,
	}, e.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var actualExporter GradeLogEntryExporter
	f, err := NewGradeLogEntryFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg, _ := f.Create()
	agg.ExportTo(&actualExporter)
	assert.Equal(t, GradeLogEntryExporter{
		EndorsedId:      seedwork.NewUint64Exporter(f.EndorsedId),
		EndorsedVersion: f.EndorsedVersion,
		AssignedGrade:   seedwork.NewUint8Exporter(f.AssignedGrade),
		CreatedAt:       f.CreatedAt,
	}, actualExporter)
}
