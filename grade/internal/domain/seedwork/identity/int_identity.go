package identity

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

func NewIntIdentity(value uint) (IntIdentity, error) {
	return IntIdentity{value: value}, nil
}

type IntIdentity struct {
	value uint
}

func (id IntIdentity) Equal(other specification.EqualOperand) bool {
	exportableOther := other.(Accessable[uint])
	return id.value == exportableOther.Value()
}

func (id IntIdentity) Export(ex exporters.ExporterSetter[uint]) {
	ex.SetState(id.value)
}

func (id IntIdentity) Value() uint {
	return id.value
}

func (id IntIdentity) String() string {
	return fmt.Sprintf("%d", id.value)
}
