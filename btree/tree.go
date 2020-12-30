package btree

import (
	"github.com/rrylee/go-algorithm/container"
)

type BTree interface {
	Search(key Key) *Node
	Insert(key Key)
	Remove(key Key)
}

type Key int

func (k Key) Compare(i interface{}) container.CompareRet {
	if i.(Key) == k {
		return container.CompareEqual
	} else if i.(Key) < k {
		return container.CompareGt
	} else {
		return container.CompareLt
	}
}

type bTree struct {
	root  *Node
	hot   *Node
	size  int
	order int //阶数
}

func (b *bTree) Search(key Key) *Node {
	v := b.root
	b.hot = nil
	for v != nil {
		i := v.Keys.Search(key)
		if i >= 0 && v.Keys.Get(i).(Key) == key {
			return v
		}
		b.hot = v
		v = v.Children.Get(i + 1).(*Node)
	}
	return nil
}

func (b bTree) Insert(key Key) {
	panic("implement me")
}

func (b bTree) Remove(key Key) {
	panic("implement me")
}
