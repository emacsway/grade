package values

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func NewUrl(val string) (Url, error) {
	return Url(val), nil
}

type Url string

func (u Url) Export(ex exporters.ExporterSetter[string]) {
	ex.SetState(string(u))
}
