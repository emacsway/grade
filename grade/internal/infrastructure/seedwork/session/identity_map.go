package session

import "errors"

type IsolationLevel int

type NonexistentObject struct{}

var ErrNonexistentObject = errors.New("")

const (
	ReadUncommittedLevel IsolationLevel = iota
	ReadCommittedLevel
	RepeatableReadsLevel
	SerializableLevel
)

type IsolationStrategy[K any] interface {
	add(key K, value any) error
	get(key K) (any, error)
	has(key K) (bool, error)
}

type SerializableStrategyImpl[K comparable] struct {
	identityMap *IdentityMapImpl[K]
}

func (s *SerializableStrategyImpl[K]) add(key K, value any) error {
	if value == nil {
		value = NonexistentObject{}
	}

	s.identityMap.doAdd(key, value)
	return nil
}

func (s *SerializableStrategyImpl[K]) get(key K) (any, error) {

	entity := s.identityMap.doGet(key)
	if _, ok := entity.(NonexistentObject); ok || entity == nil {
		return nil, ErrNonexistentObject
	}

	return entity, nil
}

func (s *SerializableStrategyImpl[K]) has(key K) (bool, error) {
	return s.identityMap.doHas(key), nil
}

type IdentityMapImpl[K comparable] struct {
	alive    map[K]any
	strategy IsolationStrategy[K]
}

func NewIdentityMap[K comparable](isolation IsolationLevel) IdentityMap[K] {
	identity := &IdentityMapImpl[K]{
		alive: map[K]any{},
	}

	identity.SetIsolationLevel(isolation)
	return identity
}

func (i *IdentityMapImpl[K]) Get(key K) (any, error) {
	return i.strategy.get(key)
}

func (i *IdentityMapImpl[K]) Add(key K, value any) error {
	return i.strategy.add(key, value)
}

func (i *IdentityMapImpl[K]) Has(key K) (bool, error) {
	return i.strategy.has(key)
}

func (i *IdentityMapImpl[K]) Clear() {
	i.alive = map[K]any{}
}

func (i *IdentityMapImpl[K]) Remove(key K) {
	if _, found := i.alive[key]; !found {
		return
	}

	delete(i.alive, key)
}

func (i *IdentityMapImpl[K]) SetIsolationLevel(isolation IsolationLevel) {
	switch isolation {
	case SerializableLevel:
		i.strategy = &SerializableStrategyImpl[K]{i}
	}
}

func (i *IdentityMapImpl[K]) doAdd(key K, value any) {
	i.alive[key] = value
}

func (i *IdentityMapImpl[K]) doGet(key K) any {
	return i.alive[key]
}

func (i *IdentityMapImpl[K]) doHas(key K) bool {
	_, found := i.alive[key]
	return found
}
