package recognizer

import (
	"testing"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
)

func TestRecognizerExport(t *testing.T) {
	f := NewRecognizerFakeFactory()
	agg, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, RecognizerState{
		Id:                        f.Id,
		MemberId:                  f.MemberId,
		Grade:                     f.Grade,
		AvailableEndorsementCount: recognizer.YearlyEndorsementCount,
		PendingEndorsementCount:   0,
		Version:                   0,
		CreatedAt:                 f.CreatedAt,
	}, agg.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var actualExporter RecognizerExporter
	f := NewRecognizerFakeFactory()
	agg, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.ExportTo(&actualExporter)
	assert.Equal(t, RecognizerExporter{
		Id:                        seedwork.NewUint64Exporter(f.Id),
		MemberId:                  seedwork.NewUint64Exporter(f.MemberId),
		Grade:                     seedwork.NewUint8Exporter(f.Grade),
		AvailableEndorsementCount: seedwork.NewUint8Exporter(recognizer.YearlyEndorsementCount),
		PendingEndorsementCount:   seedwork.NewUint8Exporter(0),
		Version:                   0,
		CreatedAt:                 f.CreatedAt,
	}, actualExporter)
}
