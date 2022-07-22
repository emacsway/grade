package recognizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecognizerExport(t *testing.T) {
	f := NewRecognizerFakeFactory()
	agg, _ := f.Create()
	assert.Equal(t, f.Export(), agg.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var expectedExporter, actualExporter RecognizerExporter
	f := NewRecognizerFakeFactory()
	f.ExportTo(&expectedExporter)
	agg, _ := f.Create()
	agg.ExportTo(&actualExporter)
	assert.Equal(t, expectedExporter, actualExporter)
}
