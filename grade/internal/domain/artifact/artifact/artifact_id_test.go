package artifact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewArtifactId(value)
	assert.Equal(t, value, id.CreateMemento())
}
