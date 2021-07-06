package vector

import (
	"strconv"

	"github.com/rrylee/go-algorithm/container"
)

type (
	Vector interface {
		Size() int
		Empty() bool
		// Search 返回不大于 key 的最后一个值
		Search(key container.Item) int
		Insert(pos int, v interface{})
		Remove(pos int) interface{}
		InsertTail(v interface{})
		Get(i int) interface{}
		Set(pos int, item interface{})
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

func (vector vector) Set(pos int, item interface{}) {
	vector.list[pos] = item
}

func (vector vector) Search(key container.Item) int {
	start, end := 0, len(vector.list)
	for start < end {
		mi := (start + end) >> 1
		if key.Compare(vector.list[mi]) == container.CompareLt {
			end = mi // [start, mi)
		} else {
			start = mi+1 // [mi+1, end)
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

func (vector *vector) Remove(pos int) (ret interface{}) {
	ret = vector.list[pos]
	vector.list = append(vector.list[:pos], vector.list[pos+1:]...)
	return
}

func (v vector) Size() int {
	return len(v.list)
}

func (v vector) Empty() bool {
	return len(v.list) == 0
}
