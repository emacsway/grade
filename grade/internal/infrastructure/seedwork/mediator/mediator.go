package mediator

import (
	"errors"
	"reflect"
	"sync"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/disposable"
)

var (
	ErrNonCallableHandler         = errors.New("")
	ErrUnsuitableHandlerSignature = errors.New("")

	plug = struct{}{}
)

type RefUntypedMediator struct {
	hLock    sync.RWMutex
	handlers map[string]reflect.Value

	sLock       sync.RWMutex
	subscribers map[string]map[reflect.Value]struct{}
}

func NewRefUntypedMediator() *RefUntypedMediator {
	return &RefUntypedMediator{
		hLock:    sync.RWMutex{},
		handlers: map[string]reflect.Value{},

		sLock:       sync.RWMutex{},
		subscribers: map[string]map[reflect.Value]struct{}{},
	}
}

func (m *RefUntypedMediator) Send(command any) {
	m.hLock.RLock()
	defer m.hLock.RUnlock()

	valueType := getValueType(command)
	if handler, found := m.handlers[valueType]; found {
		call(handler, command)
	}
}

func (m *RefUntypedMediator) Register(command any, handler any) (disposable.Disposable, error) {
	m.hLock.Lock()
	defer m.hLock.Unlock()

	if err := compareWithHandlerSignature(command, handler); err != nil {
		return nil, err
	}

	commandType := getValueType(command)
	m.handlers[commandType] = reflect.ValueOf(handler)

	return disposable.NewDisposable(func() {
		m.Unregister(command)
	}), nil
}

func (m *RefUntypedMediator) Unregister(command any) {
	m.hLock.Lock()
	defer m.hLock.Unlock()

	commandType := getValueType(command)
	delete(m.handlers, commandType)
}

func (m *RefUntypedMediator) Subscribe(event any, handler any) (disposable.Disposable, error) {
	m.sLock.Lock()
	defer m.sLock.Unlock()

	if err := compareWithHandlerSignature(event, handler); err != nil {
		return nil, err
	}

	valueType := getValueType(event)
	if _, found := m.subscribers[valueType]; !found {
		m.subscribers[valueType] = map[reflect.Value]struct{}{}
	}

	handlerValue := reflect.ValueOf(handler)
	m.subscribers[valueType][handlerValue] = plug

	return disposable.NewDisposable(func() {
		m.Unsubscribe(event, handler)
	}), nil
}

func (m *RefUntypedMediator) Unsubscribe(event any, handler any) {
	m.sLock.Lock()
	defer m.sLock.Unlock()

	eventType := getValueType(event)
	handlerValue := reflect.ValueOf(handler)

	delete(m.subscribers[eventType], handlerValue)
}

func (m *RefUntypedMediator) Publish(event any) {
	m.sLock.RLock()
	defer m.sLock.RUnlock()

	eventType := getValueType(event)
	for handler, _ := range m.subscribers[eventType] {
		call(handler, event)
	}
}

func call(callable reflect.Value, args ...any) {
	in := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}

	callable.Call(in)
}

func compareWithHandlerSignature(initiator any, handler any) error {
	handlerType := reflect.TypeOf(handler)
	if handlerType.Kind() != reflect.Func {
		return ErrNonCallableHandler
	}

	if handlerType.NumIn() < 1 {
		return ErrUnsuitableHandlerSignature
	}

	initiatorType := reflect.TypeOf(initiator)
	if handlerType.In(0) != initiatorType {
		return ErrUnsuitableHandlerSignature
	}

	return nil
}

func getValueType(t any) string {
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v.Type().String()
}
