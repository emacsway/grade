package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberIdConstructor(t *testing.T) {
	val := uint(3)
	id, err := NewMemberId(val)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, val, id.Value())
}
