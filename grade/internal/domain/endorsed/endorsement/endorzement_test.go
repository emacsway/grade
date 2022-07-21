package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEndorsementCreateMemento(t *testing.T) {
	recognizerId, _ := recognizer.NewRecognizerId(1)
	recognizerGrade, _ := shared.NewGrade(2)
	recognizerVersion := uint(3)
	endorsedId, _ := endorsed.NewEndorsedId(4)
	endorsedGrade, _ := shared.NewGrade(1)
	endorsedVersion := uint(5)
	artifactId, _ := artifact.NewArtifactId(6)
	createdAt := time.Now()
	agg, _ := NewEndorsement(
		recognizerId, recognizerGrade, recognizerVersion,
		endorsedId, endorsedGrade, endorsedVersion,
		artifactId, createdAt,
	)
	assert.Equal(t, EndorsementMemento{
		1, 2, 3, 4, 1, 5, 6, createdAt,
	}, agg.CreateMemento())
}
