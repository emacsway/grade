package seedwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityEqualsTrue(t *testing.T) {
	id, _ := NewIdentity[uint64](3)
	otherId, _ := NewIdentity[uint64](3)
	assert.True(t, id.Equals(otherId))
}

func TestIdentityEqualsFalse(t *testing.T) {
	id, _ := NewIdentity[uint64](3)
	otherId, _ := NewIdentity[uint64](4)
	assert.False(t, id.Equals(otherId))
}

func TestIdentityExport(t *testing.T) {
	var ex Uint64Exporter
	id, _ := NewIdentity[uint64](3)
	id.Export(&ex)
	assert.Equal(t, uint64(ex), id.Value())
}
