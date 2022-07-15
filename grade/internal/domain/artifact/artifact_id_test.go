package artifact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactIdEqualsTrue(t *testing.T) {
	id := NewArtifactId(3)
	otherId := NewArtifactId(3)
	assert.True(t, id.Equals(otherId))
}
