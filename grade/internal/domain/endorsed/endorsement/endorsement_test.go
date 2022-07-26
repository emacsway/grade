package endorsement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
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

func TestEndorsementIsEndorsedBy(t *testing.T) {
	cases := []struct {
		RecogniserId     uint64
		ArtifactId       uint64
		TestRecogniserId uint64
		TestArtifactId   uint64
		ExpectedResult   bool
	}{
		{1, 2, 1, 2, true},
		{1, 2, 2, 2, false},
		{1, 2, 1, 1, false},
	}
	f := NewEndorsementFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f.RecognizerId = c.RecogniserId
			rId, err := recognizer.NewRecognizerId(c.TestRecogniserId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			f.ArtifactId = c.ArtifactId
			aId, err := artifact.NewArtifactId(c.TestArtifactId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			ent, err := f.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			assert.Equal(t, c.ExpectedResult, ent.IsEndorsedBy(rId, aId))
		})
	}
}

func TestEndorsementWeight(t *testing.T) {
	f := NewEndorsementFakeFactory()
	for i := uint8(0); i <= shared.MaxGradeValue; i++ {
		for j := i; j <= shared.MaxGradeValue; j++ {
			t.Run(fmt.Sprintf("Case i=%d j=%d", i, j), func(t *testing.T) {
				var expectedWeight Weight = 1
				f.RecognizerGrade = j
				f.EndorsedGrade = i
				ent, err := f.Create()
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				if j > i {
					expectedWeight = 2
				}
				assert.Equal(t, expectedWeight, ent.GetWeight())
			})
		}
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
