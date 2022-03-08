package sort

import "sync"

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums))
}

// [l, r)
func mergeSort(nums []int, l, r int) {
	if l >= r-1 {
		return
	}

	wg := sync.WaitGroup{}
	mid := l + (r-l)/2
	go func() {
		wg.Add(1)
		mergeSort(nums, l, mid)
		wg.Done()
	}()

	mergeSort(nums, mid, r)
	wg.Wait()

	s1 := l
	s2 := mid
	for s1 < mid && s2 < r {
		if nums[s1] < nums[s2] {
			s1++
		} else {
			left := s1
			mid := s2
			right := s2 + 1
			for right < r && nums[right] <= nums[left] {
				right++
			}
			// [left, mid, right)
			rotate(nums, left, mid, right)
			s1 = right - mid + s1
			s2 = right
		}
	}
}

func rotate(nums []int, left int, mid int, right int) {
	reverse(nums, left, mid)
	reverse(nums, mid, right)
	reverse(nums, left, right)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right-1] = nums[right-1], nums[left]
		left++
		right--
	}
}
