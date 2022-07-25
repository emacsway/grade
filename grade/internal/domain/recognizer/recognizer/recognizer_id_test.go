package recognizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecognizerIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewRecognizerId(value)
	assert.Equal(t, value, id.Export())
}