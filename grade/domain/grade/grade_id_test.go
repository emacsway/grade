package grade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradeIdEqualsTrue(t *testing.T) {
	gradeId := NewGradeId(3)
	otherId := NewGradeId(3)
	assert.True(t, gradeId.Equals(otherId))
}

func TestGradeIdEqualsFalse(t *testing.T) {
	gradeId := NewGradeId(3)
	otherId := NewGradeId(4)
	assert.False(t, gradeId.Equals(otherId))
}
