package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorsedConstructor(t *testing.T) {
	grade, _ := shared.NewGrade(0)
	agg, _ := NewEndorsed(
		endorsed.NewEndorsedId(uint64(3)),
		external.NewUserId(uint64(2)),
		grade,
		[]endorsement.Endorsement{},
		1,
	)
	assert.Equal(t, 1, agg.Id.GetValue())
}
