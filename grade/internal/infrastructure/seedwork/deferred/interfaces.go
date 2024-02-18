package deferred

type DeferredCallback[T interface{}] func(T) error

type Deferred[T interface{}] interface {
	Then(DeferredCallback[T], DeferredCallback[error]) Deferred[any]
}
