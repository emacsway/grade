package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndorsedCreateMemento(t *testing.T) {
	f := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	f.AddReceivedEndorsement(rf)
	f.AddReceivedEndorsement(rf)
	agg, _ := f.Create()
	assert.Equal(t, f.CreateMemento(), agg.CreateMemento())
}
