package collections

import (
	"container/list"
	"errors"
)

var (
	ErrKeyDoesNotContains = errors.New("")
)

type CachedMap[K comparable, V any] struct {
	m map[K]V
}

func NewCachedMap[K comparable, V any]() CachedMap[K, V] {
	return CachedMap[K, V]{
		m: map[K]V{},
	}
}

func (m CachedMap[K, V]) Add(key K, value V) {
	m.m[key] = value
}

func (m CachedMap[K, V]) Get(key K) (value V, err error) {
	if value, found := m.m[key]; found {
		return value, nil
	}

	return value, ErrKeyDoesNotContains
}

func (m CachedMap[K, V]) Remove(key K) {
	delete(m.m, key)
}

func (m CachedMap[K, V]) Has(key K) bool {
	_, found := m.m[key]
	return found
}

type item[K comparable, V any] struct {
	key   K
	value V
}

type ReplacingMap[K comparable, V any] struct {
	items map[K]*list.Element
	order *list.List
	size  uint
}

func NewReplacingMap[K comparable, V any](size uint) ReplacingMap[K, V] {
	return ReplacingMap[K, V]{
		items: make(map[K]*list.Element, size),
		order: list.New(),
		size:  size,
	}
}

func (m *ReplacingMap[K, V]) Add(key K, value V) {
	element := m.order.PushBack(item[K, V]{key, value})
	m.items[key] = element

	if (uint)(len(m.items)) > m.size {
		element = m.order.Front()
		m.order.Remove(element)
		delete(m.items, element.Value.(item[K, V]).key)
	}
}

func (m *ReplacingMap[K, V]) Get(key K) (value V, err error) {
	if element, found := m.items[key]; found {
		return element.Value.(item[K, V]).value, nil
	}

	return value, ErrKeyDoesNotContains
}

func (m *ReplacingMap[K, V]) Touch(key K) {
	element := m.items[key]
	m.order.MoveToBack(element)
}

func (m *ReplacingMap[K, V]) Remove(key K) {
	element, found := m.items[key]
	if !found {
		return
	}

	delete(m.items, key)
	m.order.Remove(element)
}

func (m *ReplacingMap[K, V]) Has(key K) bool {
	_, found := m.items[key]
	return found
}

func (m *ReplacingMap[K, V]) SetSize(size uint) {
	m.size = size
}
