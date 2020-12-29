package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListQueue(t *testing.T) {
	q := NewListQueue()
	assert.Equal(t, q.Size(), 0)

	_, ok := q.Dequeue()
	assert.False(t, ok)

	q.Enqueue(1)
	assert.Equal(t, 1, q.Size())
	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, v.(int))
	assert.Equal(t, 0, q.Size())

	q.Enqueue(2)
	q.Enqueue(3)

	v, ok = q.Dequeue()
	assert.Equal(t, 2, v.(int))

	v, ok = q.Dequeue()
	assert.Equal(t, 3, v.(int))

	_, ok = q.Dequeue()
	assert.False(t, ok)
}
