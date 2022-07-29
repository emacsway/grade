package recognizer

import (
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndorsementCountConstructor(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedError error
	}{
		{uint8(0), nil},
		{YearlyEndorsementCount / 2, nil},
		{YearlyEndorsementCount, nil},
		{YearlyEndorsementCount + 1, ErrInvalidEndorsementCount},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, err := NewEndorsementCount(c.Arg)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.Arg, uint8(g))
			}
		})
	}
}

func TestEndorsementCountHasAvailable(t *testing.T) {
	cases := []struct {
		Arg            uint8
		ExpectedResult bool
	}{
		{uint8(0), false},
		{YearlyEndorsementCount / 2, true},
		{YearlyEndorsementCount, true},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := NewEndorsementCount(c.Arg)
			r := g.HasAvailable()
			assert.Equal(t, c.ExpectedResult, r)
		})
	}
}

func TestEndorsementCountNext(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedValue uint8
		ExpectedError error
	}{
		{uint8(0), uint8(0), ErrInvalidEndorsementCount},
		{YearlyEndorsementCount / 2, YearlyEndorsementCount/2 - 1, nil},
		{YearlyEndorsementCount, YearlyEndorsementCount - 1, nil},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := NewEndorsementCount(c.Arg)
			n, err := g.Decrease()
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.ExpectedValue, uint8(n))
			}
		})
	}
}

func TestEndorsementCountExportTo(t *testing.T) {
	var ex seedwork.Uint8Exporter
	c, _ := NewEndorsementCount(1)
	c.ExportTo(&ex)
	assert.Equal(t, uint8(ex), c.Export())
}
