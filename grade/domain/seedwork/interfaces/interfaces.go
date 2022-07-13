package interfaces

type Identity[T any] interface {
	Equals(Identity[T]) bool
	GetValue() T
}
