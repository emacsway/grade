package identity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntIdentityEqual(t *testing.T) {
	cases := []struct {
		Left           uint
		Right          uint
		ExpectedResult bool
	}{
		{0, 1, false},
		{0, 0, false},
		{1, 1, true},
		{3, 3, true},
		{3, 4, false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			left, err := NewIntIdentity(c.Left)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			right, err := NewIntIdentity(c.Right)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			result := left.Equal(right)
			assert.Equal(t, c.ExpectedResult, result)
		})
	}
}

func TestIntIdentityExport(t *testing.T) {
	var ex uint
	val := uint(3)
	id, _ := NewIntIdentity(val)
	id.Export(func(v uint) { ex = v })
	assert.Equal(t, val, ex)
}
