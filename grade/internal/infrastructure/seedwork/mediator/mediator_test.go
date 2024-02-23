package mediator

import (
	"testing"

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
				handler := func(t Event) {
					counter++
				}

				_, err := m.Subscribe(Event{}, handler)
				assert.NoError(t, err)

				m.Publish(Event{})
				assert.Equal(t, 1, counter)
			},
		},

		{
			name: "test_unsubscribe",
			assertion: func(t *testing.T, m *RefUntypedMediator) {

				times := 0
				handler := func(e Event) {
					times++
				}

				times2 := 0
				handler2 := func(e Event) {
					times2++
				}

				_, err := m.Subscribe(Event{}, handler)
				assert.NoError(t, err)

				_, err = m.Subscribe(Event{}, handler2)
				assert.NoError(t, err)

				m.Unsubscribe(Event{}, handler)
				m.Publish(Event{})

				assert.Equal(t, 0, times)
				assert.Equal(t, 1, times2)
			},
		},

		{
			name: "test_disposable_event",
			assertion: func(t *testing.T, m *RefUntypedMediator) {

				times := 0
				handler := func(e Event) {
					times++
				}

				times2 := 0
				handler2 := func(e Event) {
					times2++
				}

				disposable, err := m.Subscribe(Event{}, handler)
				assert.NoError(t, err)

				_, err = m.Subscribe(Event{}, handler2)
				assert.NoError(t, err)

				disposable.Dispose()
				m.Publish(Event{})

				assert.Equal(t, 0, times)
				assert.Equal(t, 1, times2)
			},
		},

		{
			name: "test_send",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(e Command) {
					times++
				}

				_, err := m.Register(Command{}, handler)
				assert.NoError(t, err)

				m.Send(Command{})

				assert.Equal(t, 1, times)
			},
		},

		{
			name: "test_unregister",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(e Command) {
					times++
				}

				_, err := m.Register(Command{}, handler)
				assert.NoError(t, err)
				m.Unregister(Command{})

				m.Send(Command{})

				assert.Equal(t, 0, times)
			},
		},

		{
			name: "test_disposable_command",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(e Command) {
					times++
				}

				disposable, err := m.Register(Command{}, handler)
				assert.NoError(t, err)

				disposable.Dispose()
				m.Send(Command{})

				assert.Equal(t, 0, times)
			},
		},

		{
			name: "test_unsuitable_params_type_handler",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(e Event) {
					times++
				}

				_, err := m.Register(Command{}, handler)
				assert.Equal(t, ErrUnsuitableHandlerSignature, err)
			},
		},

		{
			name: "test_unsuitable_params_count_handler",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				times := 0
				handler := func(e Event, smt any) {
					times++
				}

				_, err := m.Register(Command{}, handler)
				assert.Equal(t, ErrUnsuitableHandlerSignature, err)
			},
		},

		{
			name: "test_unsuitable_typeof_handler",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				handler := 1

				_, err := m.Subscribe(Command{}, handler)
				assert.Equal(t, ErrNonCallableHandler, err)
			},
		},

		{
			name: "test_unsuitable_void_handler",
			assertion: func(t *testing.T, m *RefUntypedMediator) {
				handler := func() {}

				_, err := m.Register(Command{}, handler)
				assert.Equal(t, ErrUnsuitableHandlerSignature, err)
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
