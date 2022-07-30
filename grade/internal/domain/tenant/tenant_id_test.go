package tenant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTenantIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewTenantId(value)
	assert.Equal(t, value, id.Export())
}
