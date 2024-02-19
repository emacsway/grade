package disposable

func NewDisposable(callback Callback) DisposableImp {
	return DisposableImp{callback: callback}
}

type DisposableImp struct {
	callback Callback
}

func (d DisposableImp) Dispose() {
	d.callback()
}

func (d DisposableImp) Add(other Disposable) Disposable {
	return NewCompositeDisposable(d, other)
}

func NewCompositeDisposable(disposables ...Disposable) CompositeDisposableImp {
	return CompositeDisposableImp{delegates: disposables}
}

type CompositeDisposableImp struct {
	delegates []Disposable
}

func (d CompositeDisposableImp) Dispose() {
	for i := range d.delegates {
		d.delegates[i].Dispose()
	}
}

func (d CompositeDisposableImp) Add(other Disposable) Disposable {
	return NewCompositeDisposable(append(d.delegates, other)...)
}
