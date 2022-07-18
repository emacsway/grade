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
	id, _ := endorsed.NewEndorsedId(uint64(1))
	userId, _ := external.NewUserId(uint64(2))
	grade, _ := shared.NewGrade(0)
	agg, _ := NewEndorsed(id, userId, grade, []endorsement.Endorsement{}, 1)
	assert.Equal(t, id, agg.id)
}
