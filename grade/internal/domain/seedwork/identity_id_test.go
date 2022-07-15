package seedwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorsedIdEqualsTrue(t *testing.T) {
	id := NewIdentity(3)
	otherId := NewIdentity(3)
	assert.True(t, id.Equals(otherId))
}

func TestEndorsedIdEqualsFalse(t *testing.T) {
	id := NewIdentity(3)
	otherId := NewIdentity(4)
	assert.False(t, id.Equals(otherId))
}
