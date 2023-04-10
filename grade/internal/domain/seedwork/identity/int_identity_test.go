package identity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestIntIdentityEqual(t *testing.T) {
	cases := []struct {
		Left           uint
		Right          uint
		ExpectedResult error
	}{}
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
	var ex exporters.UintExporter
	val := uint(3)
	id, _ := NewIntIdentity(val)
	id.Export(&ex)
	assert.Equal(t, val, id.Value())
}
