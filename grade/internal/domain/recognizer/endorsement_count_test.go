package recognizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestEndorsementCountConstructor(t *testing.T) {
	cases := []struct {
		Arg           uint
		ExpectedError error
	}{
		{uint(0), nil},
		{YearlyEndorsementCount / 2, nil},
		{YearlyEndorsementCount, nil},
		{YearlyEndorsementCount + 1, ErrInvalidEndorsementCount},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, err := NewEndorsementCount(c.Arg)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.Arg, uint(g))
			}
		})
	}
}

func TestEndorsementCountHasAvailable(t *testing.T) {
	cases := []struct {
		Arg            uint
		ExpectedResult bool
	}{
		{uint(0), false},
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
		Arg           uint
		ExpectedValue uint
		ExpectedError error
	}{
		{uint(0), uint(0), ErrInvalidEndorsementCount},
		{YearlyEndorsementCount / 2, YearlyEndorsementCount/2 - 1, nil},
		{YearlyEndorsementCount, YearlyEndorsementCount - 1, nil},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := NewEndorsementCount(c.Arg)
			n, err := g.Decrease()
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.ExpectedValue, uint(n))
			}
		})
	}
}

func TestEndorsementCountExport(t *testing.T) {
	var ex exporters.UintExporter
	c, _ := NewEndorsementCount(1)
	c.Export(&ex)
	assert.Equal(t, uint(c), uint(ex))
}
