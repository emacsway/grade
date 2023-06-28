package values

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestFullNameExport(t *testing.T) {
	var actualExporter FullNameExporter
	f := NewFullNameFaker()
	cid, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.Export(&actualExporter)
	assert.Equal(t, FullNameExporter{
		FirstName: exporters.StringExporter(f.FirstName),
		LastName:  exporters.StringExporter(f.LastName),
	}, actualExporter)
}
