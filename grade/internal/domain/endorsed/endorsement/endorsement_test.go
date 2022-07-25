package endorsement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorsementConstructor(t *testing.T) {
	cases := []struct {
		RecognizerGrade uint8
		EndorsedGrade   uint8
		ExpectedError   error
	}{
		{0, 0, nil},
		{1, 0, nil},
		{0, 1, ErrLowerGradeEndorses},
	}
	f := NewEndorsementFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f.RecognizerGrade = c.RecognizerGrade
			f.EndorsedGrade = c.EndorsedGrade
			e, err := f.Create()
			assert.Equal(t, f.RecognizerGrade, c.RecognizerGrade)
			assert.ErrorIs(t, err, c.ExpectedError)
			if err == nil {
				assert.Equal(t, f.Export(), e.Export())
			}
		})
	}
}

func TestEndorsementExport(t *testing.T) {
	f := NewEndorsementFakeFactory()
	e, _ := f.Create()
	assert.Equal(t, f.Export(), e.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var expectedExporter, actualExporter EndorsementExporter
	f := NewEndorsementFakeFactory()
	f.ExportTo(&expectedExporter)
	agg, _ := f.Create()
	agg.ExportTo(&actualExporter)
	assert.Equal(t, expectedExporter, actualExporter)
}
