package values

import (
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

func NewName(val string) (Name, error) {
	return Name(val), nil
}

type Name string

func (n Name) Export(ex exporters.ExporterSetter[string]) {
	ex.SetState(string(n))
}
