package artifact

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func TestTenantArtifactIdEqual(t *testing.T) {
	t1 := uuid.ParseSilent("63e8d541-af30-4593-a8ac-761dc268926d")
	m1 := uuid.ParseSilent("7c4435dc-6b5d-4628-a1f8-596dde6704b6")
	t2 := uuid.ParseSilent("e2d9fcaa-565e-4295-9142-bd69e26581cf")
	m2 := uuid.ParseSilent("c8858e26-6bc6-4775-a3bd-084773216b79")
	cases := []struct {
		TenantId        uuid.Uuid
		ArtifactId      uuid.Uuid
		OtherTenantId   uuid.Uuid
		OtherArtifactId uuid.Uuid
		ExpectedResult  bool
	}{
		{t1, m2, t1, m2, true},
		{t1, m1, t1, m2, false},
		{t2, m2, t1, m2, false},
		{t1, m2, t1, m1, false},
		{t1, m2, t2, m2, false},
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
	f := NewTenantArtifactIdFakeFactory()
	cid, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.Export(&actualExporter)
	assert.Equal(t, TenantArtifactIdExporter{
		TenantId:   exporters.UuidExporter(f.TenantId),
		ArtifactId: exporters.UuidExporter(f.ArtifactId),
	}, actualExporter)
}
