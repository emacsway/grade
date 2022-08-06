package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func NewDescription(val string) (Description, error) {
	return Description(val), nil
}

type Description string

func (d Description) Export(ex exporters.ExporterSetter[string]) {
	ex.SetState(string(d))
}
