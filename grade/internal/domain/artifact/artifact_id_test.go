package artifact

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func TestArtifactIdConstructor(t *testing.T) {
	val := uuid.ParseSilent("63e8d541-af30-4593-a8ac-761dc268926d")
	id, err := NewArtifactId(val)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, val, id.Value())
}
