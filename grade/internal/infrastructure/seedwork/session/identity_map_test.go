package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Model struct {
	pk int
}

func TestIdentityMap(t *testing.T) {

	tests := []struct {
		name     string
		testCase func(t *testing.T)
	}{
		{
			name: "Test IdentityMap with Serializable Level",
			testCase: func(t *testing.T) {
				idMap := NewIdentityMap[int](SerializableLevel)

				model := Model{3}

				err := idMap.Add(model.pk, model)
				assert.NoError(t, err)

				exists, err := idMap.Has(model.pk)
				assert.NoError(t, err)
				assert.Equal(t, true, exists)

				result, err := idMap.Get(model.pk)
				assert.NoError(t, err)

				assert.Equal(t, model, result)

				_, err = idMap.Get(10)
				assert.Equal(t, ErrNonexistentObject, err)

				err = idMap.Add(10, nil)
				assert.NoError(t, err)

				_, err = idMap.Get(10)
				assert.Equal(t, ErrNonexistentObject, err)
			},
		},

		{
			name: "Test IdentityMap object removing",
			testCase: func(t *testing.T) {
				idMap := NewIdentityMap[int](SerializableLevel)

				_ = idMap.Add(3, Model{3})
				_ = idMap.Add(5, Model{5})

				idMap.Remove(3)
				idMap.Remove(1) // remove non-exists

				_, err := idMap.Get(3)
				assert.Equal(t, ErrNonexistentObject, err)

				_, err = idMap.Get(5)
				assert.NoError(t, err)
			},
		},

		{
			name: "Test IdentityMap clearing",
			testCase: func(t *testing.T) {
				idMap := NewIdentityMap[int](SerializableLevel)

				models := []Model{
					Model{3},
					Model{5},
					Model{15},
				}

				for _, model := range models {
					err := idMap.Add(model.pk, model)
					assert.NoError(t, err)
				}

				for _, model := range models {
					_, err := idMap.Get(model.pk)
					assert.NoError(t, err)
				}

				idMap.Clear()

				for _, model := range models {
					_, err := idMap.Get(model.pk)
					assert.Equal(t, ErrNonexistentObject, err)
				}
			},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.testCase(t)
		})
	}
}
