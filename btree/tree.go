package btree

import (
	"github.com/rrylee/go-algorithm/container"
)

type BTree interface {
	Search(key container.Compare) *Node
	Insert(key container.Compare)
	Remove(key container.Compare)
}

type bTree struct {
	root  *Node
	hot   *Node
	size  int
	order int
}

func (b *bTree) Search(key container.Compare) *Node {
	v := b.root
	b.hot = nil
	for v != nil {
		i := v.Keys.Search(key)
		if i >= 0 {
			return v
		}
		b.hot = v
		v = v.Children.Get(i+1).(*Node)
	}
	return nil
}

func (b bTree) Insert(key container.Compare) {
	panic("implement me")
}

func (b bTree) Remove(key container.Compare) {
	panic("implement me")
}
