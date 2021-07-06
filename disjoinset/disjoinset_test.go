package disjoinset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisjoinSet(t *testing.T) {
	/*
		1	2
		   / \
		  3   4
		 /    /
		5 -- 6
	*/
	ds := NewDjSet(7)
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
		{4, 6},
		{5, 6},
	}

	cycle := false
	for _, edge := range edges {
		if !ds.Merge(edge[0], edge[1]) {
			cycle = true
			break
		}
	}

	assert.True(t, cycle)
}
