package endorsement

import (
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorsementConstructor(t *testing.T) {
	cases := []struct {
		RecogniserId    uint64
		RecognizerGrade uint8
		EndorsedId      uint64
		EndorsedGrade   uint8
		ExpectedError   error
	}{
		{1, 0, 2, 0, nil},
		{1, 1, 2, 0, nil},
		{1, 0, 2, 1, ErrLowerGradeEndorses},
		{1, 0, 1, 0, ErrEndorsementOneself},
	}
	f := NewEndorsementFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f.RecognizerId = c.RecogniserId
			f.RecognizerGrade = c.RecognizerGrade
			f.EndorsedId = c.EndorsedId
			f.EndorsedGrade = c.EndorsedGrade
			_, err := f.Create()
			assert.Equal(t, f.RecognizerGrade, c.RecognizerGrade)
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestEndorsementExport(t *testing.T) {
	f := NewEndorsementFakeFactory()
	e, _ := f.Create()
	assert.Equal(t, EndorsementState{
		RecognizerId:      f.RecognizerId,
		RecognizerGrade:   f.RecognizerGrade,
		RecognizerVersion: f.RecognizerVersion,
		EndorsedId:        f.EndorsedId,
		EndorsedGrade:     f.EndorsedGrade,
		EndorsedVersion:   f.EndorsedVersion,
		ArtifactId:        f.ArtifactId,
		CreatedAt:         f.CreatedAt,
	}, e.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var actualExporter EndorsementExporter
	f := NewEndorsementFakeFactory()
	agg, _ := f.Create()
	agg.ExportTo(&actualExporter)
	assert.Equal(t, EndorsementExporter{
		RecognizerId:      seedwork.NewUint64Exporter(f.RecognizerId),
		RecognizerGrade:   seedwork.NewUint8Exporter(f.RecognizerGrade),
		RecognizerVersion: f.RecognizerVersion,
		EndorsedId:        seedwork.NewUint64Exporter(f.EndorsedId),
		EndorsedGrade:     seedwork.NewUint8Exporter(f.EndorsedGrade),
		EndorsedVersion:   f.EndorsedVersion,
		ArtifactId:        seedwork.NewUint64Exporter(f.ArtifactId),
		CreatedAt:         f.CreatedAt,
	}, actualExporter)
}
