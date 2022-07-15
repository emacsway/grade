package endorsed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorsedIdEqualsTrue(t *testing.T) {
	id := NewEndorsedId(3)
	otherId := NewEndorsedId(3)
	assert.True(t, id.Equals(otherId))
}

func TestEndorsedIdEqualsFalse(t *testing.T) {
	id := NewEndorsedId(3)
	otherId := NewEndorsedId(4)
	assert.False(t, id.Equals(otherId))
}
