package identity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func TestIdentityEqual(t *testing.T) {
	cases := []struct {
		Left           uuid.Uuid
		Right          uuid.Uuid
		ExpectedResult error
	}{}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			left, err := NewUuidIdentity(c.Left)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			right, err := NewUuidIdentity(c.Right)
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
	var ex seedwork.UuidExporter
	val := uuid.ParseSilent("63e8d541-af30-4593-a8ac-761dc268926d")
	id, _ := NewUuidIdentity(val)
	id.Export(&ex)
	assert.Equal(t, val, id.Value())
}
