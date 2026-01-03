package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		FirstName: f.FirstName,
		LastName:  f.LastName,
	}, actualExporter)
}
