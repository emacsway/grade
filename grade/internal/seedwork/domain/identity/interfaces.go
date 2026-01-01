package identity

type Accessible[T any] interface {
	Value() T
}
