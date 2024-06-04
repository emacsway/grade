package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCachedMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		assertion func(t *testing.T, cm CachedMap[string, string])
	}{
		{
			name: "must contains a set key",
			assertion: func(t *testing.T, cm CachedMap[string, string]) {
				cm.Add("1", "1")
				assert.Equal(t, true, cm.Has("1"))
			},
		},

		{
			name: "must not contain a removed key",
			assertion: func(t *testing.T, cm CachedMap[string, string]) {
				cm.Add("1", "1")
				cm.Remove("1")

				assert.Equal(t, false, cm.Has("1"))
			},
		},

		{
			name: "the key can be reused",
			assertion: func(t *testing.T, cm CachedMap[string, string]) {
				cm.Add("1", "1")
				cm.Remove("1")
				cm.Add("1", "2")

				val, _ := cm.Get("1")
				assert.Equal(t, "2", val)
			},
		},

		{
			name: "trying to access an unset key",
			assertion: func(t *testing.T, cm CachedMap[string, string]) {
				_, err := cm.Get("1")
				assert.Equal(t, ErrKeyDoesNotContains, err)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cm := NewCachedMap[string, string]()
			tt.assertion(t, cm)
		})
	}
}

func TestReplacingMapMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		assertion func(t *testing.T, cm ReplacingMap[string, string])
	}{
		{
			name: "must contains a set key",
			assertion: func(t *testing.T, cm ReplacingMap[string, string]) {
				cm.Add("1", "1")
				assert.Equal(t, true, cm.Has("1"))
			},
		},

		{
			name: "must not contain a removed key",
			assertion: func(t *testing.T, cm ReplacingMap[string, string]) {
				cm.Add("1", "1")
				cm.Remove("1")

				assert.Equal(t, false, cm.Has("1"))
			},
		},

		{
			name: "the key can be reused",
			assertion: func(t *testing.T, cm ReplacingMap[string, string]) {
				cm.Add("1", "1")
				cm.Remove("1")
				cm.Add("1", "2")

				val, _ := cm.Get("1")
				assert.Equal(t, "2", val)
			},
		},

		{
			name: "trying to access an unset key",
			assertion: func(t *testing.T, cm ReplacingMap[string, string]) {
				_, err := cm.Get("1")
				assert.Equal(t, ErrKeyDoesNotContains, err)
			},
		},

		{
			name: "trying to remove an unset key",
			assertion: func(t *testing.T, cm ReplacingMap[string, string]) {
				cm.Remove("1")
				cm.Add("1", "1")

				val, _ := cm.Get("1")
				assert.Equal(t, "1", val)
				assert.Equal(t, true, cm.Has("1"))
			},
		},

		{
			name: "touched key must be in map",
			assertion: func(t *testing.T, cm ReplacingMap[string, string]) {

				cm.SetSize(2)

				cm.Add("1", "1")
				cm.Add("2", "2")

				cm.Touch("1")
				cm.Add("3", "3")

				assert.Equal(t, true, cm.Has("1"))
				assert.Equal(t, true, cm.Has("3"))

				assert.Equal(t, false, cm.Has("2"))
			},
		},

		{
			name: "trying to access to replaced key",
			assertion: func(t *testing.T, cm ReplacingMap[string, string]) {

				cm.SetSize(2)

				cm.Add("1", "1")
				cm.Add("2", "2")
				cm.Add("3", "3")

				_, err := cm.Get("1")
				assert.Equal(t, ErrKeyDoesNotContains, err)
				assert.Equal(t, false, cm.Has("1"))

				assert.Equal(t, true, cm.Has("2"))
				assert.Equal(t, true, cm.Has("3"))

				val, _ := cm.Get("2")
				assert.Equal(t, "2", val)

				val, _ = cm.Get("3")
				assert.Equal(t, "3", val)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cm := NewReplacingMap[string, string](3)
			tt.assertion(t, cm)
		})
	}
}
