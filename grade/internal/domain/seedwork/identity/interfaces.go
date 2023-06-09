package identity

type Accessable[T any] interface {
	Value() T
}
