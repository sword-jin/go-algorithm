package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/rrylee/go-algorithm/container"
	"github.com/stretchr/testify/assert"
)

type intVal int

// Compare intVal 为最小堆
func (i intVal) Compare(i2 interface{}) container.CompareRet {
	if i > i2.(intVal) {
		return container.CompareGt
	} else if i < i2.(intVal) {
		return container.CompareLt
	} else {
		return container.CompareEqual
	}
}

func TestHeap(t *testing.T) {
	heap := NewHeap()
	assert.True(t, heap.Empty())

	heap.Insert(intVal(10))
	assert.Equal(t, heap.Peek(), intVal(10))
	heap.Insert(intVal(13))
	assert.Equal(t, heap.Peek(), intVal(10))
	heap.Insert(intVal(3))
	assert.Equal(t, heap.Peek(), intVal(3))
	heap.Insert(intVal(12))
	assert.Equal(t, heap.Peek(), intVal(3))
	heap.Insert(intVal(1))
	assert.Equal(t, heap.Peek(), intVal(1))

	v, ok := heap.Pop()
	assert.True(t, ok)
	assert.Equal(t, v, intVal(1))
	v, ok = heap.Pop()
	assert.True(t, ok)
	assert.Equal(t, v, intVal(3))
	v, ok = heap.Pop()
	assert.True(t, ok)
	assert.Equal(t, v, intVal(10))
	v, ok = heap.Pop()
	assert.True(t, ok)
	assert.Equal(t, v, intVal(12))
	v, ok = heap.Pop()
	assert.True(t, ok)
	assert.Equal(t, v, intVal(13))
}

func TestCopyFrom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var nums = []container.Item{
		intVal(rand.Int()),
		intVal(rand.Int()),
		intVal(rand.Int()),
		intVal(rand.Int()),
		intVal(rand.Int()),
	}
	heapI := CopyFrom(nums)
	heap := heapI.(*heap)
	assert.Less(t, int(heap.list[0].(intVal)), int(heap.list[1].(intVal)))
	assert.Less(t, int(heap.list[0].(intVal)), int(heap.list[2].(intVal)))
	assert.Less(t, int(heap.list[1].(intVal)), int(heap.list[3].(intVal)))
	assert.Less(t, int(heap.list[1].(intVal)), int(heap.list[4].(intVal)))
}

