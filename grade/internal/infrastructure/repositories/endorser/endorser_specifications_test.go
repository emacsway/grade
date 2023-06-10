package endorser

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestEndorserCanCompleteEndorsementSpecification(t *testing.T) {
	sp := EndorserCanCompleteEndorsementSpecification{}
	sql, params, err := sp.Compile()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(
		t,
		"endorser.available_endorsement_count != $1 AND "+
			"endorser.pending_endorsement_count != $2 AND "+
			"endorser.available_endorsement_count >= endorser.pending_endorsement_count",
		sql)
	assert.Equal(t, []driver.Valuer{
		exporters.UintExporter(0),
		exporters.UintExporter(0),
	}, params)
}
