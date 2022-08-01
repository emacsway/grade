package expertisearea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpertiseAreaIdConstructor(t *testing.T) {
	var value uint64 = 3
	id, _ := NewExpertiseAreaId(value)
	assert.Equal(t, value, id.Value())
}
