package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberIdEqualsTrue(t *testing.T) {
	memberId := NewMemberId(3)
	otherId := NewMemberId(3)
	assert.True(t, memberId.Equals(otherId))
}

func TestMemberIdEqualsFalse(t *testing.T) {
	memberId := NewMemberId(3)
	otherId := NewMemberId(4)
	assert.False(t, memberId.Equals(otherId))
}
