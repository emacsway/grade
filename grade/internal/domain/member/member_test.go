package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
)

func TestMemberExport(t *testing.T) {
	var actualExporter MemberExporter
	f := NewMemberFaker()
	agg, err := f.Create(nil)
	require.NoError(t, err)
	agg.Export(&actualExporter)
	var expectedStatus uint8
	f.Status.Export(func(v uint8) { expectedStatus = v })
	assert.Equal(t, MemberExporter{
		Id:     values.NewMemberIdExporter(f.Id.TenantId, f.Id.MemberId),
		Status: expectedStatus,
		FullName: values.FullNameExporter{
			FirstName: f.FullName.FirstName,
			LastName:  f.FullName.LastName,
		},
		CreatedAt: f.CreatedAt,
	}, actualExporter)
}
