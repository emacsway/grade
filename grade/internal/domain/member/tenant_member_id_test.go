package member

import (
	"testing"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
)

func TestTenantMemberIdExport(t *testing.T) {
	cid, err := NewTenantMemberId(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, TenantMemberIdState{
		TenantId: 1,
		MemberId: 2,
	}, cid.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var actualExporter TenantMemberIdExporter
	cid, err := NewTenantMemberId(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.ExportTo(&actualExporter)
	assert.Equal(t, TenantMemberIdExporter{
		TenantId: seedwork.NewUint64Exporter(1),
		MemberId: seedwork.NewUint64Exporter(2),
	}, actualExporter)
}
