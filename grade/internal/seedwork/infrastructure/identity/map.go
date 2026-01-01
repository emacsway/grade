package identity

import (
	"errors"

	"github.com/emacsway/grade/grade/pkg/collections"
)

var (
	ErrObjectAlreadyWatched = errors.New("")
	ErrObjectNotFound       = errors.New("")
)

type IdentityMap[K comparable, V any] struct {
	manageable collections.ReplacingMap[K, V]
	isolation  IsolationStrategy[K, V]
}

func NewIdentityMap[K comparable, V any](size uint) *IdentityMap[K, V] {
	manageable := collections.NewReplacingMap[K, V](size)
	isolation := serializableStrategy[K, V]{manageable: manageable}

	return &IdentityMap[K, V]{
		manageable: manageable,
		isolation:  &isolation,
	}
}

func (im *IdentityMap[K, V]) Add(key K, object V) (bool, error) {
	if err := im.isolation.add(key, object); err != nil {
		return false, err
	}

	return true, nil
}

func (im *IdentityMap[K, V]) Get(key K) (object V, err error) {
	return im.isolation.get(key)
}

func (im *IdentityMap[K, V]) Has(key K) bool {
	return im.isolation.has(key)
}

func (im *IdentityMap[K, V]) SetSize(size uint) {
	im.manageable.SetSize(size)
}

func (im *IdentityMap[K, V]) SetIsolationLevel(level IsolationLevel) {

	switch level {
	case ReadUncommitted:
		im.isolation = &readUncommittedStrategy[K, V]{manageable: im.manageable}
	case RepeatableReads:
		im.isolation = &repeatableReadsStrategy[K, V]{manageable: im.manageable}
	case Serializable:
		im.isolation = &serializableStrategy[K, V]{manageable: im.manageable}
	case ReadCommitted:
		im.isolation = &readCommittedStrategy[K, V]{manageable: im.manageable}
	default:
		im.isolation = &serializableStrategy[K, V]{manageable: im.manageable}
	}
}
