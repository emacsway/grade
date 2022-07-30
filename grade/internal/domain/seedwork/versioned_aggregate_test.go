package seedwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionedAggregateConstructor(t *testing.T) {
	var value uint = 3
	va, _ := NewVersionedAggregate(value)
	assert.Equal(t, value, va.GetVersion())
	va.IncreaseVersion()
	assert.Equal(t, value+1, va.GetVersion())
}
