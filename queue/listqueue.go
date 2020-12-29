package queue

import "container/list"

type listQueue struct {
	list *list.List
}

func (l listQueue) Size() int {
	return l.list.Len()
}

func (l listQueue) Empty() bool {
	return l.Size() == 0
}

func (l listQueue) Enqueue(val interface{}) {
	l.list.PushBack(val)
}

func (l listQueue) Dequeue() (interface{}, bool) {
	if l.Empty() {
		return nil, false
	}
	e := l.list.Front()
	l.list.Remove(e)
	return e.Value, true
}

func NewListQueue() *listQueue {
	return &listQueue{list: list.New()}
}
