package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewUrl(val string) (Url, error) {
	return Url(val), nil
}

type Url string

func (u Url) Export(ex identity.ExporterSetter[string]) {
	ex.SetState(string(u))
}
