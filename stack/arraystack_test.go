package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()

	_, ok := s.Peek()
	assert.False(t, ok)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 3, v.(int))

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v.(int))

	v, ok = s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 2, v.(int))

	assert.Equal(t, s.Size(), 2)
	assert.False(t, s.Empty())
}

