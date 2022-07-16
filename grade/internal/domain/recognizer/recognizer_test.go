package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecognizerConstructor(t *testing.T) {
	id, _ := recognizer.NewRecognizerId(uint64(1))
	userId, _ := external.NewUserId(uint64(2))
	grade, _ := shared.NewGrade(0)
	count, _ := recognizer.NewAvailableEndorsementCount(uint8(20))
	agg, _ := NewRecognizer(id, userId, grade, count, 1)
	assert.Equal(t, id, agg.Id)
}
