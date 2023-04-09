package competence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestTenantCompetenceIdEqual(t *testing.T) {
	t1 := uint(10)
	m1 := uint(3)
	t2 := uint(11)
	m2 := uint(4)
	cases := []struct {
		TenantId          uint
		CompetenceId      uint
		OtherTenantId     uint
		OtherCompetenceId uint
		ExpectedResult    bool
	}{
		{t1, m2, t1, m2, true},
		{t1, m1, t1, m2, false},
		{t2, m2, t1, m2, false},
		{t1, m2, t1, m1, false},
		{t1, m2, t2, m2, false},
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
	f := NewTenantCompetenceIdFakeFactory()
	cid, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.Export(&actualExporter)
	assert.Equal(t, TenantCompetenceIdExporter{
		TenantId:     exporters.UintExporter(f.TenantId),
		CompetenceId: exporters.UintExporter(f.CompetenceId),
	}, actualExporter)
}
