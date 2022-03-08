package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{-1, -2, 3, -6, 9, 4}
	MergeSort(nums)
	fmt.Println(nums)
}
