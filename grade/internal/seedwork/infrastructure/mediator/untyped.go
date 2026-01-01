package mediator

import (
	"context"
	"errors"
	"reflect"
	"sync"

	"github.com/emacsway/grade/grade/internal/seedwork/domain/disposable"
	"github.com/hashicorp/go-multierror"
)

var (
	ErrUnsuitableHandlerSignature = errors.New("passed handler has unsuitable signature")
)

type Handler func(ctx context.Context, command any) (any, error)

func AsUntyped[T any](handler func(ctx context.Context, command T) (any, error)) Handler {
	return func(ctx context.Context, command any) (any, error) {
		if typedCommand, ok := command.(T); ok {
			return handler(ctx, typedCommand)
		}

		return nil, ErrUnsuitableHandlerSignature
	}
}

type RefUntypedMediator struct {
	hLock    sync.RWMutex
	handlers map[reflect.Type]Handler

	sLock       sync.RWMutex
	subscribers map[reflect.Type]map[reflect.Value]Handler

	pLock sync.RWMutex
	pipes []func(next Handler) Handler
}

func NewRefUntypedMediator() *RefUntypedMediator {
	return &RefUntypedMediator{
		hLock:    sync.RWMutex{},
		handlers: map[reflect.Type]Handler{},

		sLock:       sync.RWMutex{},
		subscribers: map[reflect.Type]map[reflect.Value]Handler{},

		pLock: sync.RWMutex{},
	}
}

func (m *RefUntypedMediator) AddPipe(pipe func(next Handler) Handler) {
	m.pLock.Lock()
	defer m.pLock.Unlock()

	m.pipes = append(m.pipes, pipe)
}

func (m *RefUntypedMediator) executeWithPipeline(handler Handler, ctx context.Context, command any) (any, error) {
	m.pLock.RLock()
	defer m.pLock.RUnlock()

	current := func(ctx context.Context, command any) (any, error) {
		return handler(ctx, command)
	}

	for ixd := range m.pipes {
		reverse := len(m.pipes) - 1 - ixd
		current = m.pipes[reverse](current)
	}

	return current(ctx, command)
}

func (m *RefUntypedMediator) Send(ctx context.Context, command any) (any, error) {
	m.hLock.RLock()
	defer m.hLock.RUnlock()

	commandType := reflect.TypeOf(command)
	if handler, found := m.handlers[commandType]; found {
		return m.executeWithPipeline(handler, ctx, command)
	}

	return nil, nil
}

func (m *RefUntypedMediator) Register(command any, handler Handler) disposable.Disposable {
	m.hLock.Lock()
	defer m.hLock.Unlock()

	commandType := reflect.TypeOf(command)
	m.handlers[commandType] = handler

	return disposable.NewDisposable(func() {
		m.Unregister(command)
	})
}

func (m *RefUntypedMediator) Unregister(command any) {
	m.hLock.Lock()
	defer m.hLock.Unlock()

	commandType := reflect.TypeOf(command)
	delete(m.handlers, commandType)
}

func (m *RefUntypedMediator) Subscribe(event any, handler Handler) disposable.Disposable {
	m.sLock.Lock()
	defer m.sLock.Unlock()

	valueType := reflect.TypeOf(event)
	if _, found := m.subscribers[valueType]; !found {
		m.subscribers[valueType] = map[reflect.Value]Handler{}
	}

	handlerValue := reflect.ValueOf(handler)
	m.subscribers[valueType][handlerValue] = handler

	return disposable.NewDisposable(func() {
		m.Unsubscribe(event, handler)
	})
}

func (m *RefUntypedMediator) Unsubscribe(event any, handler Handler) {
	m.sLock.Lock()
	defer m.sLock.Unlock()

	eventType := reflect.TypeOf(event)
	handlerValue := reflect.ValueOf(handler)

	delete(m.subscribers[eventType], handlerValue)
}

func (m *RefUntypedMediator) Publish(ctx context.Context, event any) error {
	m.sLock.RLock()
	defer m.sLock.RUnlock()

	var errs error
	eventType := reflect.TypeOf(event)
	for _, handler := range m.subscribers[eventType] {
		_, err := handler(ctx, event)
		errs = multierror.Append(errs, err)
	}

	return errs
}
