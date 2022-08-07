package recognizer

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestRecognizerCanCompleteEndorsementSpecification(t *testing.T) {
	sp := RecognizerCanCompleteEndorsementSpecification{}
	sql, params, err := sp.Compile()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(
		t,
		"recognizer.available_endorsement_count != ? AND "+
			"recognizer.pending_endorsement_count != ? AND "+
			"recognizer.available_endorsement_count >= recognizer.pending_endorsement_count",
		sql)
	assert.Equal(t, []driver.Valuer{
		exporters.UintExporter(0),
		exporters.UintExporter(0),
	}, params)
}
