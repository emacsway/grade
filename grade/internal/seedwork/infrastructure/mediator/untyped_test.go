package mediator

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

type (
	Event struct {
		name string
	}
	Command struct {
		name string
	}
)

func TestMediator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		assertion func(t *testing.T, m *RefUntypedMediator)
	}{
		{
			name: "test_publish",

			assertion: func(t *testing.T, m *RefUntypedMediator) {
				counter := 0
				handler := func(ctx context.Context, t any) (any, error) {
					counter++

					return nil, nil
				}

				m.Subscribe(Event{}, handler)

				_ = m.Publish(context.Background(), Event{})
				assert.Equal(t, 1, counter)
			},
		},

		{
			name: "test_unsubscribe",
			assertion: func(t *testing.T, m *RefUntypedMediator) {

				times := 0
				handler := func(ctx context.Context, e any) (any, error) {
					times++
					return nil, nil
				}

				times2 := 0
				handler2 := func(ctx context.Context, e any) (any, error) {
					times2++
					return nil, nil
				}

				times3 := 0
				handler3 := AsUntyped(func(ctx context.Context, e any) (any, error) {
					times3++
					return nil, nil
				})

				m.Subscribe(Event{}, handler)
				m.Subscribe(Event{}, handler2)
				m.Subscribe(Event{}, handler3)

				m.Unsubscribe(Event{}, handler)
				_ = m.Publish(context.Background(), Event{})

				assert.Equal(t, 0, times)
				assert.Equal(t, 1, times2)
				assert.Equal(t, 1, times3)

				m.Unsubscribe(Event{}, handler3)
				_ = m.Publish(context.Background(), Event{})

				assert.Equal(t, 0, times)
				assert.Equal(t, 2, times2)
				assert.Equal(t, 1, times3)
			},
		},

		{
			name: "test_disposable_event",
			assertion: func(t *testing.T, m *RefUntypedMediator) {

				times := 0
				handler := func(ctx context.Context, e any) (any, error) {
					times++
					return nil, nil
				}

				times2 := 0
				handler2 := func(ctx context.Context, e any) (any, error) {
					times2++
					return nil, nil
				}

				disposable := m.Subscribe(Event{}, handler)
				m.Subscribe(Event{}, handler2)

				disposable.Dispose()
				_ = m.Publish(context.Background(), Event{})

				assert.Equal(t, 0, times)
				assert.Equal(t, 1, times2)
			},
		},

		{
			name: "test_send",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(ctx context.Context, e any) (any, error) {
					times++
					return nil, nil
				}

				m.Register(Command{}, handler)
				_, _ = m.Send(context.Background(), Command{})

				assert.Equal(t, 1, times)
			},
		},

		{
			name: "test_unregister",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(ctx context.Context, e any) (any, error) {
					times++

					return nil, nil
				}

				m.Register(Command{}, handler)
				m.Unregister(Command{})

				_, _ = m.Send(context.Background(), Command{})

				assert.Equal(t, 0, times)
			},
		},

		{
			name: "test_disposable_command",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(ctx context.Context, e any) (any, error) {
					times++

					return nil, nil
				}

				disposable := m.Register(Command{}, handler)
				disposable.Dispose()

				_, _ = m.Send(context.Background(), Command{})

				assert.Equal(t, 0, times)
			},
		},

		{
			name: "test_unsuitable_params_type_handler",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(ctx context.Context, e int) (any, error) {
					times++

					return nil, nil
				}

				m.Register(Command{}, AsUntyped(handler))

				_, err := m.Send(context.Background(), Command{})
				assert.Equal(t, ErrUnsuitableHandlerSignature, err)
			},
		},

		{
			name: "test_returning_errors",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				handlerError := errors.New("")

				handler := func(ctx context.Context, e any) (any, error) {
					return nil, handlerError
				}

				handler2 := func(ctx context.Context, e any) (any, error) {
					return nil, handlerError
				}

				handler3 := func(ctx context.Context, e Event) (any, error) {
					return nil, handlerError
				}

				m.Register(Command{}, handler)

				m.Subscribe(Event{}, handler2)
				m.Subscribe(Event{}, AsUntyped(handler3))

				var errs error
				errs = multierror.Append(errs, handlerError, handlerError)
				assert.Equal(t, errs, m.Publish(context.Background(), Event{}))

				_, err := m.Send(context.Background(), Command{})
				assert.Equal(t, handlerError, err)
			},
		},
		{
			name: "test_execute_pipeline",
			assertion: func(t *testing.T, m *RefUntypedMediator) {

				m.AddPipe(func(next Handler) Handler {

					return AsUntyped[Command](func(ctx context.Context, command Command) (any, error) {
						return next(ctx, Command{name: command.name + "1"})
					})

				})

				m.AddPipe(func(next Handler) Handler {

					return AsUntyped[Command](func(ctx context.Context, command Command) (any, error) {
						return next(ctx, Command{name: command.name + "2"})
					})

				})

				handler := func(ctx context.Context, e any) (any, error) {
					command := e.(Command)
					return Command{name: command.name + "3"}, nil
				}

				m.Register(Command{}, handler)

				res, err := m.Send(context.Background(), Command{})
				assert.NoError(t, err)

				assert.Equal(t, Command{name: "123"}, res)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			m := NewRefUntypedMediator()
			tt.assertion(t, m)
		})
	}
}

func asd() {
	m := NewRefUntypedMediator()

	handler3 := func(ctx context.Context, e Event) (any, error) {
		return nil, nil
	}

	m.Subscribe(Event{}, AsUntyped(handler3))
}
