package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUuid(t *testing.T) {
	id := NewUuid()
	assert.IsType(t, id, Uuid{})
}

func TestParse(t *testing.T) {
	val := "63e8d541-af30-4593-a8ac-761dc268926d"
	id, err := Parse(val)
	assert.Nil(t, err)
	assert.IsType(t, val, id.String())
}
