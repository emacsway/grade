package competence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompetenceIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewCompetenceId(value)
	assert.Equal(t, value, id.Value())
}
