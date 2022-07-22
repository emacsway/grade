package shared

import (
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeConstructor(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedError error
	}{
		{uint8(0), nil},
		{maxGradeValue / 2, nil},
		{maxGradeValue, nil},
		{maxGradeValue + 1, ErrInvalidGrade},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, err := NewGrade(c.Arg)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.Arg, uint8(g))
			}
		})
	}
}

func TestGradeHasNext(t *testing.T) {
	cases := []struct {
		Arg            uint8
		ExpectedResult bool
	}{
		{uint8(0), true},
		{maxGradeValue / 2, true},
		{maxGradeValue, false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := NewGrade(c.Arg)
			r := g.HasNext()
			assert.Equal(t, c.ExpectedResult, r)
		})
	}
}

func TestGradeNext(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedValue uint8
		ExpectedError error
	}{
		{uint8(0), uint8(1), nil},
		{maxGradeValue / 2, maxGradeValue/2 + 1, nil},
		{maxGradeValue, uint8(0), ErrInvalidGrade},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := NewGrade(c.Arg)
			n, err := g.Next()
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.ExpectedValue, uint8(n))
			}
		})
	}
}

func TestGradeSetState(t *testing.T) {
	g, _ := NewGrade(1)
	g.Import(2)
	assert.Equal(t, uint8(2), uint8(g))
}

func TestGradeExportTo(t *testing.T) {
	var ex seedwork.Uint8Exporter
	g, _ := NewGrade(1)
	g.ExportTo(&ex)
	assert.Equal(t, uint8(ex), g.Export())
}
