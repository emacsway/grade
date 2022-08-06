package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewDescription(val string) (Description, error) {
	return Description(val), nil
}

type Description string

func (d Description) Export(ex identity.ExporterSetter[string]) {
	ex.SetState(string(d))
}
