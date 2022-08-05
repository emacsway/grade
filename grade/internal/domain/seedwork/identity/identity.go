package identity

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

type ExporterSetter[T any] interface {
	SetState(T)
}

type Exportable[T any] interface {
	Export(ExporterSetter[T])
}

// alternative approach:

type Accessable[T any] interface {
	Value() T
}

type Equaler interface {
	Equal(Equaler) bool
}

func NewUuidIdentity(value uuid.Uuid) (UuidIdentity, error) {
	return UuidIdentity{value: value}, nil
}

type UuidIdentity struct {
	value uuid.Uuid
}

func (id UuidIdentity) Equal(other Equaler) bool {
	exportableOther := other.(Accessable[uuid.Uuid])
	return id.value == exportableOther.Value()
}

func (id UuidIdentity) Export(ex ExporterSetter[uuid.Uuid]) {
	ex.SetState(id.value)
}

func (id UuidIdentity) Value() uuid.Uuid {
	return id.value
}

func (id UuidIdentity) String() string {
	return fmt.Sprintf("%v", id.value)
}
