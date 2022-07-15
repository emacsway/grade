package endorsed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorsedIdEqualsTrue(t *testing.T) {
	memberId := NewEndorsedId(3)
	otherId := NewEndorsedId(3)
	assert.True(t, memberId.Equals(otherId))
}

func TestEndorsedIdEqualsFalse(t *testing.T) {
	memberId := NewEndorsedId(3)
	otherId := NewEndorsedId(4)
	assert.False(t, memberId.Equals(otherId))
}
