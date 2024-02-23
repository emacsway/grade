package mediator

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/disposable"
)

type Mediator interface {
	Register(commandType any, handler any) (disposable.Disposable, error)
	Unregister(commandType any)
	Send(command any)

	Subscribe(eventType any, handler any) (disposable.Disposable, error)
	Unsubscribe(eventType any, handler any)
	Publish(event any)
}
