package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector(t *testing.T) {
	v := NewVector(5)
	v.Insert(0, 2)
	v.Insert(1, 3)
	v.Insert(1, 4)
	v.Insert(0, 5)

	assert.Equal(t, 4, v.Size())

	v.Remove(0)
	assert.Equal(t, 3, v.Size())
	assert.Equal(t, v.list, []interface{}{2, 4, 3})
}
