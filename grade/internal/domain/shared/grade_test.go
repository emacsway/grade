package shared

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeConstructor(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedError error
	}{
		{uint8(0), nil},
		{uint8(1), nil},
		{uint8(2), nil},
		{uint8(3), nil},
		{uint8(4), nil},
		{uint8(5), nil},
		{uint8(6), ErrInvalidGrade},
	}
	for _, c := range cases {
		g, err := NewGrade(c.Arg)
		assert.Equal(t, c.ExpectedError, err)
		if err == nil {
			assert.Equal(t, c.Arg, uint8(g))
		}
	}
}

func TestGradeHasNext(t *testing.T) {
	cases := []struct {
		Arg            uint8
		ExpectedResult bool
	}{
		{uint8(0), true},
		{uint8(1), true},
		{uint8(2), true},
		{uint8(3), true},
		{uint8(4), true},
		{uint8(5), false},
	}
	for _, c := range cases {
		g, _ := NewGrade(c.Arg)
		r := g.HasNext()
		assert.Equal(t, c.ExpectedResult, r)
	}
}

func TestGradeNext(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedValue uint8
		ExpectedError error
	}{
		{uint8(0), uint8(1), nil},
		{uint8(1), uint8(2), nil},
		{uint8(2), uint8(3), nil},
		{uint8(3), uint8(4), nil},
		{uint8(4), uint8(5), nil},
		{uint8(5), uint8(0), ErrInvalidGrade},
	}
	for _, c := range cases {
		g, _ := NewGrade(c.Arg)
		n, err := g.Next()
		assert.Equal(t, c.ExpectedError, err)
		if err == nil {
			assert.Equal(t, c.ExpectedValue, uint8(n))
		}
	}
}
