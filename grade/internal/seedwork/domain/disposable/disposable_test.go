package disposable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisposableDispose(t *testing.T) {
	var result bool = false
	d := NewDisposable(func() { result = true })
	d.Dispose()
	assert.True(t, result)
}

func TestDisposableAdd(t *testing.T) {
	var result1 bool = false
	var result2 bool = false
	d1 := NewDisposable(func() { result1 = true })
	d2 := NewDisposable(func() { result2 = true })
	d := d1.Add(d2)
	d.Dispose()
	assert.True(t, result1)
	assert.True(t, result2)
}

func TestCompositeDisposableAdd(t *testing.T) {
	var result1 bool = false
	var result2 bool = false
	var result3 bool = false
	d1 := NewDisposable(func() { result1 = true })
	d2 := NewDisposable(func() { result2 = true })
	d3 := NewDisposable(func() { result3 = true })
	d := d1.Add(d2).Add(d3)
	d.Dispose()
	assert.True(t, result1)
	assert.True(t, result2)
	assert.True(t, result3)
}
