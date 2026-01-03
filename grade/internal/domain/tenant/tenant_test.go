package tenant

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		Id:        f.Id,
		Name:      f.Name,
		CreatedAt: f.CreatedAt,
	}, actualExporter)
}
