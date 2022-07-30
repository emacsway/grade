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

type Equaler interface {
	Equals(Equaler) bool
}

func NewIdentity[T comparable](value T) (Identity[T, ExporterSetter[T]], error) {
	return Identity[T, ExporterSetter[T]]{value: value}, nil
}

// The way to fix issue of generics:
// https://issuemode.com/issues/golang/go/105227904

type Identity[T comparable, C ExporterSetter[T]] struct {
	value T
}

func (id Identity[T, C]) Equals(other Equaler) bool {
	typedOther := other.(Exportable[T])
	return id.value == typedOther.Export()
}

func (id Identity[T, C]) Export() T {
	return id.value
}

func (id Identity[T, C]) ExportTo(ex C) {
	ex.SetState(id.value)
}
