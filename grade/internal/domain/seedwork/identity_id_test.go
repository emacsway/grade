package seedwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityEqualsTrue(t *testing.T) {
	id := NewIdentity[uint64](3)
	otherId := NewIdentity[uint64](3)
	assert.True(t, id.Equals(otherId))
}

func TestIdentityEqualsFalse(t *testing.T) {
	id := NewIdentity[uint64](3)
	otherId := NewIdentity[uint64](4)
	assert.False(t, id.Equals(otherId))
}
