package session

type DeferredImp[T interface{}] struct {
	value      T
	err        error
	onSuccess  DeferredCallback[T]
	onFailure  DeferredCallback[error]
	isResolved bool
}

func (r *DeferredImp[T]) Resolve(value T) error {
	r.value = value
	r.isResolved = true
	return r.doResolve()
}

func (r *DeferredImp[T]) Then(callback DeferredCallback[T]) error {
	r.onSuccess = callback
	if r.isResolved {
		return r.doResolve()
	}
	return nil
}

func (r *DeferredImp[T]) Catch(callback DeferredCallback[error]) error {
	r.onFailure = callback
	if r.isResolved && r.err != nil {
		return r.onFailure(r.err)
	}
	return nil
}

func (r *DeferredImp[T]) doResolve() error {
	if r.onSuccess != nil {
		r.err = r.onSuccess(r.value)
	}
	if r.err != nil {
		if r.onFailure != nil {
			return r.onFailure(r.err)
		} else {
			return r.err
		}
	}
	return nil
}
