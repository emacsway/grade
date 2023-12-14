package identity

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/specification"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewUuidIdentity(value uuid.Uuid) (UuidIdentity, error) {
	return UuidIdentity{value: value}, nil
}

type UuidIdentity struct {
	value uuid.Uuid
}

func (id UuidIdentity) Equal(other specification.EqualOperand) bool {
	exportableOther := other.(Accessible[uuid.Uuid])
	return id.value == exportableOther.Value()
}

func (id UuidIdentity) Export(ex exporters.ExporterSetter[uuid.Uuid]) {
	ex.SetState(id.value)
}

func (id UuidIdentity) Value() uuid.Uuid {
	return id.value
}

func (id UuidIdentity) String() string {
	return id.value.String()
}
