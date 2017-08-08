package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcurrentMultiMap(t *testing.T) {
	concurrentMap := NewConcurrentMultiMap()
	concurrentMap.Set("names", []interface{}{"Raed Shomali"})
	concurrentMap.Append("names", "Dwayne Johnson")

	ok := concurrentMap.Contains("names")
	assert.True(t, ok)

	size := concurrentMap.Size()
	assert.Equal(t, size, 1)

	isEmpty := concurrentMap.IsEmpty()
	assert.False(t, isEmpty)

	keys := concurrentMap.Keys()
	assert.Equal(t, keys, []string{"names"})

	values, ok := concurrentMap.Get("names")
	assert.True(t, ok)
	assert.Equal(t, values, []interface{}{"Raed Shomali", "Dwayne Johnson"})

	values, ok = concurrentMap.Get("unknown")
	assert.False(t, ok)
	assert.Nil(t, values)

	concurrentMap.Remove("names")
	values, ok = concurrentMap.Get("names")
	assert.False(t, ok)
	assert.Nil(t, values)

	concurrentMap.Append("names", "Raed Shomali")
	concurrentMap.Clear()

	values, ok = concurrentMap.Get("names")
	assert.False(t, ok)
	assert.Nil(t, values)

	size = concurrentMap.Size()
	assert.Equal(t, size, 0)

	isEmpty = concurrentMap.IsEmpty()
	assert.True(t, isEmpty)

	keys = concurrentMap.Keys()
	assert.Equal(t, keys, []string{})
}
