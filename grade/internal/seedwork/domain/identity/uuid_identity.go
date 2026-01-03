package identity

import (
	"github.com/emacsway/grade/grade/internal/seedwork/domain/specification"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/uuid"
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

func (id UuidIdentity) Export(ex func(uuid.Uuid)) {
	ex(id.value)
}

func (id UuidIdentity) Value() uuid.Uuid {
	return id.value
}

func (id UuidIdentity) String() string {
	return id.value.String()
}
