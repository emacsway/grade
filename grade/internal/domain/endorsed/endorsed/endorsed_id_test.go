package endorsed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorsedIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewEndorsedId(value)
	assert.Equal(t, value, id.Export())
}
