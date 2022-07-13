package domain

import (
	"github.com/emacsway/qualifying-grade/grade/domain/grade"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeIdEqualsTrue(t *testing.T) {
	gradeId := grade.NewGradeId(3)
	otherId := grade.NewGradeId(3)
	assert.True(t, gradeId.Equals(otherId))
}

func TestGradeIdEqualsFalse(t *testing.T) {
	gradeId := grade.NewGradeId(3)
	otherId := grade.NewGradeId(4)
	assert.False(t, gradeId.Equals(otherId))
}
