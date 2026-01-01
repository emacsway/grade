package disposable

type Callback func()

type Disposable interface {
	Dispose()
	Add(Disposable) Disposable
}
