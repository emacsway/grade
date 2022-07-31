package shared

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func TestGradeConstructor(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedError error
	}{
		{uint8(0), nil},
		{MaxGradeValue / 2, nil},
		{MaxGradeValue, nil},
		{MaxGradeValue + 1, ErrInvalidGrade},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, err := DefaultConstructor(c.Arg)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.Arg, g.value)
			}
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
		{MaxGradeValue / 2, MaxGradeValue/2 + 1, nil},
		{MaxGradeValue, uint8(0), ErrInvalidGrade},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := DefaultConstructor(c.Arg)
			n, err := g.Next()
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.ExpectedValue, n.value)
			}
		})
	}
}

func TestGradePrevious(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedValue uint8
		ExpectedError error
	}{
		{uint8(0), uint8(0), ErrInvalidGrade},
		{MaxGradeValue / 2, MaxGradeValue/2 - 1, nil},
		{MaxGradeValue, MaxGradeValue - 1, nil},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := DefaultConstructor(c.Arg)
			p, err := g.Previous()
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.ExpectedValue, p.value)
			}
		})
	}
}

func TestGradeExport(t *testing.T) {
	var ex seedwork.Uint8Exporter
	g, _ := DefaultConstructor(1)
	g.Export(&ex)
	assert.Equal(t, g.value, uint8(ex))
}
