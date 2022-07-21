package endorsement

import (
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEndorsementConstructor(t *testing.T) {
	cases := []struct {
		RecognizerGrade uint8
		EndorsedGrade   uint8
		ExpectedError   error
	}{
		{0, 0, nil},
		{1, 0, nil},
		{0, 1, ErrHigherGradeEndorsed},
	}
	f := NewEndorsementTestFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f.RecognizerGrade = c.RecognizerGrade
			f.EndorsedGrade = c.EndorsedGrade
			e, err := f.Create()
			assert.Equal(t, f.RecognizerGrade, c.RecognizerGrade)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, f.CreateMemento(), e.CreateMemento())
			}
		})
	}
}

func TestEndorsementCreateMemento(t *testing.T) {
	f := NewEndorsementTestFactory()
	e, _ := f.Create()
	assert.Equal(t, f.CreateMemento(), e.CreateMemento())
}

func NewEndorsementTestFactory() *EndorsementTestFactory {
	return &EndorsementTestFactory{
		1, 2, 3, 4, 1, 5, 6, time.Now(),
	}
}

type EndorsementTestFactory struct {
	RecognizerId      uint64
	RecognizerGrade   uint8
	RecognizerVersion uint
	EndorsedId        uint64
	EndorsedGrade     uint8
	EndorsedVersion   uint
	ArtifactId        uint64
	CreatedAt         time.Time
}

func (f EndorsementTestFactory) Create() (Endorsement, error) {
	recognizerId, _ := recognizer.NewRecognizerId(f.RecognizerId)
	recognizerGrade, _ := shared.NewGrade(f.RecognizerGrade)
	endorsedId, _ := endorsed.NewEndorsedId(f.EndorsedId)
	endorsedGrade, _ := shared.NewGrade(f.EndorsedGrade)
	artifactId, _ := artifact.NewArtifactId(f.ArtifactId)
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		endorsedId, endorsedGrade, f.EndorsedVersion,
		artifactId, f.CreatedAt,
	)
}

func (f EndorsementTestFactory) CreateMemento() EndorsementMemento {
	return EndorsementMemento{
		f.RecognizerId, f.RecognizerGrade, f.RecognizerVersion,
		f.EndorsedId, f.EndorsedGrade, f.EndorsedVersion,
		f.ArtifactId, f.CreatedAt,
	}
}
