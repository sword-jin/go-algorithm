package backpack

import (
	"fmt"
	"strconv"
	"strings"
)

// 一维 dp
func MaxValue(total int, weight []int, value []int) int {
	if total <= 0 {
		panic("total must >0")
	}
	if len(weight) != len(value) {
		panic("length must equal")
	}

	// dp[i] 表示 i 容量的背包的最大价值
	dp := make([]int, total+1)
	for i := 0; i <= total; i++ {
		if i >= weight[0] {
			dp[i] = value[0]
		}
	}

	for i := 1; i < len(weight); i++ {
		for j := total; j >= weight[i]; j-- {
			dp[j] = max(dp[j-weight[i]]+value[i], dp[j])
		}
	}
	return dp[total]
}

// 二维 dp
func MaxValue2(total int, weight []int, value []int) int {
	if total <= 0 {
		panic("total must >0")
	}
	if len(weight) != len(value) {
		panic("length must equal")
	}

	// dp[i][j] 表示使用j空间装[0,i]个背包的最大价值
	// dp[i][j] = max(dp[i-1][j], dp[i][j - weight[i]] + value[i])
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, total+1)
	}
	// 先拿第一个物品
	for i := 0; i < total+1; i++ {
		if i >= weight[0] {
			dp[0][i] = value[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		for j := 0; j <= total; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[len(weight)-1][total]
}

func printDp(dp [][]int) {
	for i := range dp {
		ss := []string{}
		for _, n := range dp[i] {
			ss = append(ss, strconv.Itoa(n))
		}
		fmt.Println(strings.Join(ss, " "))
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
