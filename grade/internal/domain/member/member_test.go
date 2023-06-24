package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestMemberExport(t *testing.T) {
	var actualExporter MemberExporter
	f := NewMemberFakeFactory()
	agg, err := f.Create()
	require.NoError(t, err)
	agg.Export(&actualExporter)
	assert.Equal(t, MemberExporter{
		Id:     values.NewTenantMemberIdExporter(f.Id.TenantId, f.Id.MemberId),
		Status: exporters.Uint8Exporter(f.Status),
		FullName: values.FullNameExporter{
			FirstName: exporters.StringExporter(f.FullName.FirstName),
			LastName:  exporters.StringExporter(f.FullName.LastName),
		},
		CreatedAt: f.CreatedAt,
	}, actualExporter)
}
