package competence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork"
)

func TestTenantCompetenceIdEqual(t *testing.T) {
	cases := []struct {
		TenantId       uint64
		CompetenceId       uint64
		OtherTenantId  uint64
		OtherCompetenceId  uint64
		ExpectedResult bool
	}{
		{1, 2, 1, 2, true},
		{1, 1, 1, 2, false},
		{2, 2, 1, 2, false},
		{1, 2, 1, 1, false},
		{1, 2, 2, 2, false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			id, err := NewTenantCompetenceId(c.TenantId, c.CompetenceId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			otherId, err := NewTenantCompetenceId(c.OtherTenantId, c.OtherCompetenceId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r := id.Equal(otherId)
			assert.Equal(t, c.ExpectedResult, r)
		})
	}
}

func TestRecognizerExport(t *testing.T) {
	var actualExporter TenantCompetenceIdExporter
	cid, err := NewTenantCompetenceId(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.Export(&actualExporter)
	assert.Equal(t, TenantCompetenceIdExporter{
		TenantId: seedwork.Uint64Exporter(1),
		CompetenceId: seedwork.Uint64Exporter(2),
	}, actualExporter)
}
