package tenant

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func TestTenantIdConstructor(t *testing.T) {
	val, err := uuid.Parse("63e8d541-af30-4593-a8ac-761dc268926d")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	id, err := NewTenantId(val)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, val, id.Value())
}
