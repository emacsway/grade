package seedwork

type ExporterSetter[T any] interface {
	SetState(T)
}

type ExportableTo[T any] interface {
	ExportTo(ExporterSetter[T])
}

// alternative approach:

type Exportable[T any] interface {
	Export() T
}

type Identifier[T comparable] interface {
	Exportable[T]
	ExportableTo[T]
	Equals(Identifier[T]) bool
}

func NewIdentity[T comparable](value T) (Identity[T, Identifier[T], ExporterSetter[T]], error) {
	return Identity[T, Identifier[T], ExporterSetter[T]]{value: value}, nil
}

// The way to fix issue of generics:
// https://issuemode.com/issues/golang/go/105227904

type Identity[T comparable, C Identifier[T], D ExporterSetter[T]] struct {
	value T
}

func (id Identity[T, C, D]) Equals(other C) bool {
	return id.value == other.Export()
}

func (id Identity[T, C, D]) Export() T {
	return id.value
}

func (id Identity[T, C, D]) ExportTo(ex D) {
	ex.SetState(id.value)
}
