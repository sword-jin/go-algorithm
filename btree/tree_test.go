package btree

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestEmptyBTree(t *testing.T) {
	btree := NewBTree(4)
	assert.Equal(t, btree.Size(), 0)
}

func TestBTree_Insert(t *testing.T) {
	btree := NewBTree(4)
	btree.Insert(Key(12))
	assert.Equal(t, btree.Size(), 1)

	btree.Insert(Key(13))
	assert.Equal(t, btree.Size(), 2)

	assert.False(t, btree.Insert(Key(13)))

	btree.Insert(Key(3))
	assert.Equal(t, btree.Size(), 3)

	// 发生分裂
	assert.True(t, btree.Insert(Key(120)))
	spew.Dump(btree)
}
