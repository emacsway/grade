package artifact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactIdConstructor(t *testing.T) {
	val := uint(3)
	id, err := NewArtifactId(val)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, val, id.Value())
}
