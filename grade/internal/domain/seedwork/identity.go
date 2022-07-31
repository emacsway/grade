package seedwork

type ExporterSetter[T any] interface {
	SetState(T)
}

type ExportableTo[T any] interface {
	Export(ExporterSetter[T])
}

// alternative approach:

type Accessable[T any] interface {
	Value() T
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
	exportableOther := other.(Accessable[T])
	return id.value == exportableOther.Value()
}

func (id Identity[T, C]) Value() T {
	return id.value
}

func (id Identity[T, C]) Export(ex C) {
	ex.SetState(id.value)
}
