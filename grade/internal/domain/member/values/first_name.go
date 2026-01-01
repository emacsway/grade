package values

import (
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func NewFirstName(val string) (FirstName, error) {
	return FirstName(val), nil
}

type FirstName string

func (n FirstName) Export(ex exporters.ExporterSetter[string]) {
	ex.SetState(string(n))
}
