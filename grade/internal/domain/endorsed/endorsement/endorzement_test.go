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
	recognizerId, _ := recognizer.NewRecognizerId(1)
	recognizerVersion := uint(3)
	endorsedId, _ := endorsed.NewEndorsedId(4)
	endorsedVersion := uint(5)
	artifactId, _ := artifact.NewArtifactId(6)
	createdAt := time.Now()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			recognizerGrade, _ := shared.NewGrade(c.RecognizerGrade)
			endorsedGrade, _ := shared.NewGrade(c.EndorsedGrade)
			e, err := NewEndorsement(
				recognizerId, recognizerGrade, recognizerVersion,
				endorsedId, endorsedGrade, endorsedVersion,
				artifactId, createdAt,
			)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, EndorsementMemento{
					1, c.RecognizerGrade, 3, 4, c.EndorsedGrade, 5, 6, createdAt,
				}, e.CreateMemento())
			}
		})
	}
}

func TestEndorsementCreateMemento(t *testing.T) {
	recognizerId, _ := recognizer.NewRecognizerId(1)
	recognizerGrade, _ := shared.NewGrade(2)
	recognizerVersion := uint(3)
	endorsedId, _ := endorsed.NewEndorsedId(4)
	endorsedGrade, _ := shared.NewGrade(1)
	endorsedVersion := uint(5)
	artifactId, _ := artifact.NewArtifactId(6)
	createdAt := time.Now()
	e, _ := NewEndorsement(
		recognizerId, recognizerGrade, recognizerVersion,
		endorsedId, endorsedGrade, endorsedVersion,
		artifactId, createdAt,
	)
	assert.Equal(t, EndorsementMemento{
		1, 2, 3, 4, 1, 5, 6, createdAt,
	}, e.CreateMemento())
}
