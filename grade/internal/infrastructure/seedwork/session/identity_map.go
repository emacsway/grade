package session

import "errors"

type IsolationLevel int

type NonexistentObject struct{}

var (
	ErrNonexistentObject = errors.New("")
	ErrUnknownKey        = errors.New("")
)

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

type SerializableStrategy[K comparable] struct {
	identityMap *IdentityMapImpl[K]
}

func (s *SerializableStrategy[K]) add(key K, value any) error {
	if value == nil {
		value = NonexistentObject{}
	}

	s.identityMap.doAdd(key, value)
	return nil
}

func (s *SerializableStrategy[K]) get(key K) (any, error) {

	object := s.identityMap.doGet(key)
	if _, ok := object.(NonexistentObject); ok || object == nil {
		return nil, ErrNonexistentObject
	}

	return object, nil
}

func (s *SerializableStrategy[K]) has(key K) (bool, error) {
	return s.identityMap.doHas(key), nil
}

type ReadUncommittedStrategy[K comparable] struct {
	identityMap *IdentityMapImpl[K]
}

func (s *ReadUncommittedStrategy[K]) add(key K, value any) error {
	return nil
}

func (s *ReadUncommittedStrategy[K]) get(key K) (any, error) {
	return nil, ErrUnknownKey
}

func (s *ReadUncommittedStrategy[K]) has(key K) (bool, error) {
	return false, nil
}

type ReadCommittedStrategy[K comparable] struct {
	identityMap *IdentityMapImpl[K]
}

func (s *ReadCommittedStrategy[K]) add(key K, value any) error {
	return nil
}

func (s *ReadCommittedStrategy[K]) get(key K) (any, error) {
	return nil, ErrUnknownKey
}

func (s *ReadCommittedStrategy[K]) has(key K) (bool, error) {
	return false, nil
}

type RepeatableReadsStrategy[K comparable] struct {
	identityMap *IdentityMapImpl[K]
}

func (s *RepeatableReadsStrategy[K]) add(key K, value any) error {
	if value != nil {
		s.identityMap.doAdd(key, value)
	}

	return nil
}

func (s *RepeatableReadsStrategy[K]) get(key K) (any, error) {
	object := s.identityMap.doGet(key)
	if _, ok := object.(NonexistentObject); ok || object == nil {
		return nil, ErrNonexistentObject
	}

	return object, nil
}

func (s *RepeatableReadsStrategy[K]) has(key K) (bool, error) {
	if !s.identityMap.doHas(key) {
		return false, ErrUnknownKey
	}

	object := s.identityMap.doGet(key)
	_, ok := object.(NonexistentObject)

	return ok, nil
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
		i.strategy = &SerializableStrategy[K]{i}
	case ReadUncommittedLevel:
		i.strategy = &ReadUncommittedStrategy[K]{i}
	case ReadCommittedLevel:
		i.strategy = &ReadCommittedStrategy[K]{i}
	case RepeatableReadsLevel:
		i.strategy = &RepeatableReadsStrategy[K]{i}
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
