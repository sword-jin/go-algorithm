package vector

import (
	"github.com/rrylee/go-algorithm/container"
)

type (
	Vector interface {
		Size() int
		Empty() bool
		// Search 返回不大于 key 的最后一个值
		Search(key container.Compare) int
		Insert(pos int, v interface{})
		Delete(pos int)
		InsertTail(v interface{})
		Get(i int) interface{}
	}

	vector struct {
		list []interface{}
	}
)

func NewVector(l int) *vector {
	return &vector{
		list: make([]interface{}, 0, l),
	}
}

func (vector vector) Get(i int) interface{} {
	return vector.list[i]
}

func (vector vector) Search(key container.Compare) int {
	start, end := 0, len(vector.list)
	for start < end {
		mi := (start + end) >> 1
		if key.Compare(vector.list[mi]) == container.CompareLt {
			start = mi
		} else {
			end = mi + 1
		}
	}
	return start - 1
}

func (vector *vector) Insert(pos int, v interface{}) {
	l := len(vector.list)
	if pos == l {
		vector.list = append(vector.list, v)
	} else {
		vector.list = append(vector.list, nil) //加一个元素
		for i := l; i > pos; i-- {
			vector.list[i] = vector.list[i-1]
		}
		vector.list[pos] = v
	}
}

func (vector *vector) InsertTail(v interface{}) {
	vector.list = append(vector.list, v)
}

func (vector *vector) Delete(pos int) {
	vector.list = append(vector.list[:pos], vector.list[pos+1:]...)
}

func (v vector) Size() int {
	return len(v.list)
}

func (v vector) Empty() bool {
	return len(v.list) == 0
}
