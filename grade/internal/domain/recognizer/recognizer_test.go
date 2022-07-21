package recognizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecognizerExport(t *testing.T) {
	f := NewRecognizerFakeFactory()
	agg, _ := f.Create()
	assert.Equal(t, f.Export(), agg.Export())
}
