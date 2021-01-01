package heap

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rrylee/go-algorithm/container"
)

type Heap interface {
	container.Container
	Insert(v container.Item)
	Peek() container.Item
	Pop() (container.Item, bool)
	Swap(i, j int)
	Filterup(start int)
	Filterdown(start, end int)
	Debug()
}

// heap 二叉堆
type heap struct {
	list []container.Item
}

func (t *heap) Debug() {
	spew.Dump(t)
}

func (t *heap) Size() int {
	return len(t.list)
}

func (t *heap) Empty() bool {
	return t.Size() == 0
}

func (t *heap) Swap(i, j int) {
	t.list[i], t.list[j] = t.list[j], t.list[i]
}

func (t *heap) Insert(v container.Item) {
	t.list = append(t.list, v)
	t.Filterup(t.Size()-1)
	return
}

func (t *heap) Peek() container.Item {
	return t.list[0]
}

func (t *heap) Pop() (item container.Item, ok bool) {
	l := len(t.list)
	if len(t.list) == 0 {
		ok = false
		return
	}

	ok = true

	item = t.list[0]
	t.list[0] = t.list[l-1]
	t.list = t.list[:l-1]

	if l == 1 {
		return
	}

	t.Filterdown(0, l-2)

	return
}

// Filterdown 向下调整（删除节点）
func (t *heap) Filterdown(start, end int) {
	tmp := t.list[start] //开始的节点
	i := 2 * start + 1 //left child 节点位置

	for i <= end {
		if i < end && t.list[i].Compare(t.list[i + 1]) == container.CompareGt { //左孩子大于右孩子
			i ++ // 选择较小的和移动节点比较
		}
		if i > end {
			break
		}
		if tmp.Compare(t.list[i]) != container.CompareGt { // >=
			break
		}
		t.list[start] = t.list[i]
		start = i
		i = 2 * i + 1
	}
	t.list[start] = tmp
}

// Filterup 向上调整（增加节点）
func (t *heap) Filterup(start int) {
	tmp := t.list[start]
	parentIndex := (start - 1) / 2

	for start > 0 {
		if t.list[parentIndex].Compare(tmp) != container.CompareGt { // <=
			break
		} else {
			t.list[start] = t.list[parentIndex]
			start = parentIndex
			parentIndex = (parentIndex - 1) / 2
		}
	}
	t.list[start] = tmp
}

func NewHeap() Heap {
	return &heap{
		list: []container.Item{},
	}
}

func CopyFrom(list []container.Item) Heap {
	heap := &heap{
		list: list,
	}

	l := len(heap.list)
	start := (l - 1)/2
	for start >= 0 {
		heap.Filterdown(start, l-1)
		start --
	}

	return heap
}
