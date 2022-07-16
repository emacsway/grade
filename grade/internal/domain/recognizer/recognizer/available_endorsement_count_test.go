package recognizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvailableEndorsementCountConstructor(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedError error
	}{
		{uint8(0), nil},
		{yearlyEndorsementCount / 2, nil},
		{yearlyEndorsementCount, nil},
		{yearlyEndorsementCount + 1, ErrInvalidAvailableEndorsementCount},
	}
	for _, c := range cases {
		g, err := NewAvailableEndorsementCount(c.Arg)
		assert.Equal(t, c.ExpectedError, err)
		if err == nil {
			assert.Equal(t, c.Arg, uint8(g))
		}
	}
}

func TestAvailableEndorsementCountHasAvailable(t *testing.T) {
	cases := []struct {
		Arg            uint8
		ExpectedResult bool
	}{
		{uint8(0), false},
		{yearlyEndorsementCount / 2, true},
		{yearlyEndorsementCount, true},
	}
	for _, c := range cases {
		g, _ := NewAvailableEndorsementCount(c.Arg)
		r := g.HasAvailable()
		assert.Equal(t, c.ExpectedResult, r)
	}
}

func TestAvailableEndorsementCountNext(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedValue uint8
		ExpectedError error
	}{
		{uint8(0), uint8(0), ErrInvalidAvailableEndorsementCount},
		{yearlyEndorsementCount / 2, yearlyEndorsementCount/2 - 1, nil},
		{yearlyEndorsementCount, yearlyEndorsementCount - 1, nil},
	}
	for _, c := range cases {
		g, _ := NewAvailableEndorsementCount(c.Arg)
		n, err := g.Decrease()
		assert.Equal(t, c.ExpectedError, err)
		if err == nil {
			assert.Equal(t, c.ExpectedValue, uint8(n))
		}
	}
}
