package tenant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTenantIdConstructor(t *testing.T) {
	val := uint(10)
	id, err := NewTenantId(val)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, val, id.Value())
}
