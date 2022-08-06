package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

func NewName(val string) (Name, error) {
	return Name(val), nil
}

type Name string

func (n Name) Export(ex identity.ExporterSetter[string]) {
	ex.SetState(string(n))
}
