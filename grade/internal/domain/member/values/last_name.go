package values

import (
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func NewLastName(val string) (LastName, error) {
	return LastName(val), nil
}

type LastName string

func (n LastName) Export(ex exporters.ExporterSetter[string]) {
	ex.SetState(string(n))
}
