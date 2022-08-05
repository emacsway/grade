package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionedAggregateConstructor(t *testing.T) {
	var value uint = 3
	va := NewVersionedAggregate(value)
	assert.Equal(t, value, va.Version())
	va.SetVersion(va.Version() + 1)
	assert.Equal(t, value+1, va.Version())
}
