package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewUserId(value)
	assert.Equal(t, value, id.Export())
}
