package member

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func TestTenantMemberIdEqual(t *testing.T) {
	cases := []struct {
		TenantId       uint64
		MemberId       uint64
		OtherTenantId  uint64
		OtherMemberId  uint64
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
			id, err := NewTenantMemberId(c.TenantId, c.MemberId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			otherId, err := NewTenantMemberId(c.OtherTenantId, c.OtherMemberId)
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
	var actualExporter TenantMemberIdExporter
	cid, err := NewTenantMemberId(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.Export(&actualExporter)
	assert.Equal(t, TenantMemberIdExporter{
		TenantId: seedwork.Uint64Exporter(1),
		MemberId: seedwork.Uint64Exporter(2),
	}, actualExporter)
}
