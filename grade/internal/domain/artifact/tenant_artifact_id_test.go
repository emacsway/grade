package artifact

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork"
)

func TestTenantArtifactIdEqual(t *testing.T) {
	cases := []struct {
		TenantId        uint64
		ArtifactId      uint64
		OtherTenantId   uint64
		OtherArtifactId uint64
		ExpectedResult  bool
	}{
		{1, 2, 1, 2, true},
		{1, 1, 1, 2, false},
		{2, 2, 1, 2, false},
		{1, 2, 1, 1, false},
		{1, 2, 2, 2, false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			id, err := NewTenantArtifactId(c.TenantId, c.ArtifactId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			otherId, err := NewTenantArtifactId(c.OtherTenantId, c.OtherArtifactId)
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
	var actualExporter TenantArtifactIdExporter
	cid, err := NewTenantArtifactId(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.Export(&actualExporter)
	assert.Equal(t, TenantArtifactIdExporter{
		TenantId:   seedwork.Uint64Exporter(1),
		ArtifactId: seedwork.Uint64Exporter(2),
	}, actualExporter)
}
