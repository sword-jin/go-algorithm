package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/rrylee/go-algorithm/container"
	"github.com/stretchr/testify/assert"
)

func TestHeapSortInt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	nums := []container.Item{
		container.IntVal(rand.Int()),
		container.IntVal(rand.Int()),
		container.IntVal(rand.Int()),
		container.IntVal(rand.Int()),
		container.IntVal(rand.Int()),
	}

	HeapSort(nums)
	assert.Greater(t, int(nums[0].(container.IntVal)), int(nums[1].(container.IntVal)))
	assert.Greater(t, int(nums[1].(container.IntVal)), int(nums[2].(container.IntVal)))
	assert.Greater(t, int(nums[2].(container.IntVal)), int(nums[3].(container.IntVal)))
	assert.Greater(t, int(nums[3].(container.IntVal)), int(nums[4].(container.IntVal)))
}
