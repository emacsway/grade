package interfaces

type Identity[T comparable] interface {
	Equals(Identity[T]) bool
	GetValue() T
}
