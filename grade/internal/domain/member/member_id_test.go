package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewMemberId(value)
	assert.Equal(t, value, id.Export())
}
