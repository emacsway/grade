package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompetenceIdConstructor(t *testing.T) {
	val := uint(3)
	id, err := NewCompetenceId(val)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, val, id.Value())
}
