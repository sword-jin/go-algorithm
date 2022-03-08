package main

import (
	"fmt"
	"sort"
	"strconv"
)

var count int

func main() {
	fmt.Println(minNumber([]int{41, 23, 87, 55, 50, 53, 18, 9, 39, 63, 35, 33, 54, 25, 26, 49, 74, 61, 32, 81, 97, 99, 38, 96, 22, 95, 35, 57, 80, 80, 16, 22, 17, 13, 89, 11, 75, 98, 57, 81, 69, 8, 10, 85, 13, 49, 66, 94, 80, 25, 13, 85, 55, 12, 87, 50, 28, 96, 80, 43, 10, 24, 88, 52, 16, 92, 61, 28, 26, 78, 28, 28, 16, 1, 56, 31, 47, 85, 27, 30, 85, 2, 30, 51, 84, 50, 3, 14, 97, 9, 91, 90, 63, 90, 92, 89, 76, 76, 67, 55}))
}

func minNumber(nums []int) string {
	totalLength := 0
	for _, num := range nums {
		if num == 0 {
			totalLength++
		} else {
			for num > 0 {
				num /= 10
				totalLength++
			}
		}
	}
	sort.Ints(nums)

	var ans []byte
	var path []byte
	used := make([]bool, len(nums))
	var backtrace func()
	backtrace = func() {
		if len(ans) != 0 && string(path) > string(ans) {
			return
		}

		if len(path) == totalLength {
			if len(ans) == 0 || string(path) < string(ans) {
				ans = make([]byte, len(path))
				copy(ans, path)
			}
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && used[i-1] == false && nums[i] == nums[i-1] {
				continue
			}
			numStr := strconv.Itoa(nums[i])
			path = append(path, numStr...)
			used[i] = true
			backtrace()
			used[i] = false
			path = path[:len(path)-len(numStr)]
		}
	}

	backtrace()
	return string(ans)
}
