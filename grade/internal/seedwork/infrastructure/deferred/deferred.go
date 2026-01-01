package deferred

import "github.com/hashicorp/go-multierror"

/**
* Simplified version of
* - https://github.com/emacsway/store/blob/devel/polyfill.js#L199
* - https://github.com/emacsway/go-promise
**/

func Noop[T interface{}](_ T) error { return nil }

type handler[T interface{}] struct {
	onSuccess DeferredCallback[T]
	onError   DeferredCallback[error]
	next      *DeferredImp[any]
}

type DeferredImp[T interface{}] struct {
	value       T
	err         error
	occurredErr error
	isResolved  bool
	isRejected  bool
	handlers    []handler[T]
}

func (d *DeferredImp[T]) Resolve(value T) {
	d.value = value
	d.isResolved = true
	for _, h := range d.handlers {
		d.resolveHandler(h)
	}
}

func (d *DeferredImp[T]) Reject(err error) {
	d.isRejected = true
	for _, h := range d.handlers {
		d.rejectHandler(h)
	}
}

func (d *DeferredImp[T]) Then(onSuccess DeferredCallback[T], onError DeferredCallback[error]) Deferred[any] {
	next := &DeferredImp[any]{}
	h := handler[T]{
		onSuccess: onSuccess,
		onError:   onError,
		next:      next,
	}
	d.handlers = append(d.handlers, h)
	if d.isResolved {
		d.resolveHandler(h)
	} else if d.isRejected {
		d.rejectHandler(h)
	}
	return next
}

func (d *DeferredImp[T]) resolveHandler(h handler[T]) {
	err := h.onSuccess(d.value)
	if err == nil {
		h.next.Resolve(true)
	} else {
		d.occurredErr = multierror.Append(d.occurredErr, err)
		h.next.Reject(err)
	}
}

func (d *DeferredImp[T]) rejectHandler(h handler[T]) {
	err := h.onError(d.err)
	if err != nil {
		d.occurredErr = multierror.Append(d.occurredErr, err)
		h.next.Reject(err)
	}
}

func (d *DeferredImp[T]) OccurredErr() error {
	err := d.occurredErr
	for _, h := range d.handlers {
		nestedErr := h.next.OccurredErr()
		if nestedErr != nil {
			err = multierror.Append(err, nestedErr)
		}
	}
	return err
}
