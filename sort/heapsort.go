package sort

import (
	"github.com/rrylee/go-algorithm/container"
	heap2 "github.com/rrylee/go-algorithm/heap"
)

func HeapSort(nums []container.Item) {
	heap := heap2.CopyFrom(nums)
	l := len(nums)
	for l > 0 {
		heap.Swap(0, l-1)
		l--
		heap.Filterdown(0, l-1)
	}
}
