package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestMemberExport(t *testing.T) {
	var actualExporter MemberExporter
	f := NewMemberFakeFactory()
	agg, err := f.Create()
	require.NoError(t, err)
	agg.Export(&actualExporter)
	assert.Equal(t, MemberExporter{
		Id:     NewTenantMemberIdExporter(f.Id.TenantId, f.Id.MemberId),
		Status: exporters.Uint8Exporter(f.Status),
		FullName: FullNameExporter{
			exporters.StringExporter(f.FullName.FirstName),
			exporters.StringExporter(f.FullName.LastName),
		},
		CreatedAt: f.CreatedAt,
	}, actualExporter)
}
