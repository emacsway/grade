package tenant

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func TestTenantExport(t *testing.T) {
	var actualExporter TenantExporter
	f := NewTenantFaker()
	agg, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, TenantExporter{
		Id:        exporters.UintExporter(f.Id),
		Name:      exporters.StringExporter(f.Name),
		CreatedAt: f.CreatedAt,
	}, actualExporter)
}
