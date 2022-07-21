package recognizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecognizerCreateMemento(t *testing.T) {
	f := NewRecognizerFakeFactory()
	agg, _ := f.Create()
	assert.Equal(t, f.CreateMemento(), agg.CreateMemento())
}
