package identity

import (
	"errors"
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
)

var (
	ErrNonTransient = errors.New("identity should be transient")
)

func NewIntIdentity(value uint) (IntIdentity, error) {
	return IntIdentity{value: value}, nil
}

func NewTransientIntIdentity() IntIdentity {
	return IntIdentity{}
}

type IntIdentity struct {
	value uint
}

func (id IntIdentity) Equal(other specification.EqualOperand) bool {
	exportableOther := other.(Accessable[uint])
	return !id.IsTransient() && id.value == exportableOther.Value()
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

func (id IntIdentity) IsTransient() bool {
	return id.value == 0
}

func (id *IntIdentity) Scan(src any) error { // Call me in InsertQuery with auto-increment PK
	if !id.IsTransient() {
		return ErrNonTransient
	}
	val, ok := src.(uint)
	if !ok {
		return errors.New("invalid type")
	}
	id.value = val
	return nil
}
