package identity

import (
	"errors"

	"github.com/emacsway/grade/grade/pkg/collections"
)

type IsolationLevel uint

const (
	ReadUncommitted IsolationLevel = iota
	ReadCommitted                  = iota
	RepeatableReads                = iota
	Serializable                   = iota
)

var (
	ErrNonexistentObject          = errors.New("")
	ErrDeniedOperationForStrategy = errors.New("")
)

type IsolationStrategy[K comparable, V any] interface {
	add(key K, object V) error
	get(key K) (V, error)
	has(key K) bool
}

type readUncommittedStrategy[K comparable, V any] struct {
	manageable collections.ReplacingMap[K, V]
}

func (r *readUncommittedStrategy[K, V]) add(key K, object V) error {
	return nil
}

func (r *readUncommittedStrategy[K, V]) get(key K) (object V, err error) {
	return object, ErrDeniedOperationForStrategy
}

func (r *readUncommittedStrategy[K, V]) has(key K) bool {
	return false
}

type readCommittedStrategy[K comparable, V any] struct {
	manageable collections.ReplacingMap[K, V]
}

func (r *readCommittedStrategy[K, V]) add(key K, object V) error {
	return nil
}

func (r *readCommittedStrategy[K, V]) get(key K) (object V, err error) {
	return object, nil
}

func (r *readCommittedStrategy[K, V]) has(key K) bool {
	return false
}

type repeatableReadsStrategy[K comparable, V any] struct {
	manageable collections.ReplacingMap[K, V]
}

func (r *repeatableReadsStrategy[K, V]) add(key K, object V) error {
	return nil
}

func (r *repeatableReadsStrategy[K, V]) get(key K) (V, error) {
	return r.manageable.Get(key)
}

func (r *repeatableReadsStrategy[K, V]) has(key K) bool {
	return r.manageable.Has(key)
}

type serializableStrategy[K comparable, V any] struct {
	manageable collections.ReplacingMap[K, V]
}

func (s *serializableStrategy[K, V]) add(key K, object V) error {
	return nil
}

func (s *serializableStrategy[K, V]) get(key K) (V, error) {
	return s.manageable.Get(key)
}

func (s *serializableStrategy[K, V]) has(key K) bool {
	return s.manageable.Has(key)
}
