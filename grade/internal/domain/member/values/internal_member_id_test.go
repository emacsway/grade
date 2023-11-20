package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalMemberIdConstructor(t *testing.T) {
	val := uint(3)
	id, err := NewInternalMemberId(val)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, val, id.Value())
}
