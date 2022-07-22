package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndorsedExport(t *testing.T) {
	ef := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	ef.AddReceivedEndorsement(rf)
	ef.AddReceivedEndorsement(rf)
	agg, _ := ef.Create()
	assert.Equal(t, ef.Export(), agg.Export())
}

func TestEndorsedExportTo(t *testing.T) {
	var expectedExporter, actualExporter EndorsedExporter
	ef := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	ef.AddReceivedEndorsement(rf)
	ef.AddReceivedEndorsement(rf)
	ef.ExportTo(&expectedExporter)
	agg, _ := ef.Create()
	agg.ExportTo(actualExporter)
	assert.Equal(t, ef.Export(), agg.Export())
}
