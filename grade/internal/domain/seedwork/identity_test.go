package seedwork

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityEqual(t *testing.T) {
	cases := []struct {
		Left           uint64
		Right          uint64
		ExpectedResult error
	}{}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			left, err := NewUint64Identity(c.Left)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			right, err := NewUint64Identity(c.Right)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			result := left.Equal(right)
			assert.Equal(t, c.ExpectedResult, result)
		})
	}
}

func TestIdentityExport(t *testing.T) {
	var ex Uint64Exporter
	id, _ := NewUint64Identity(3)
	id.Export(&ex)
	assert.Equal(t, uint64(ex), id.Value())
}
