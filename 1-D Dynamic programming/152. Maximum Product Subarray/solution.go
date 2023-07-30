package main

import "fmt"

const nInf = -1 << 33

const (
	positive = 0
	negative = 1
)

func maxProduct(nums []int) int {
	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, 2)
		for s := range memo[i] {
			memo[i][s] = nInf
		}
	}
	best := nInf
	for i := range nums {
		best = max(best, dp(nums, i, positive, memo))
	}
	return best
}

func dp(nums []int, i, sing int, memo [][]int) int {
	if i == len(nums)-1 {
		return nums[i]
	}
	if memo[i][sing] != nInf {
		return memo[i][sing]
	}
	var best int
	if (sing == positive) == (nums[i] >= 0) {
		best = dp(nums, i+1, positive, memo)
	} else {
		best = dp(nums, i+1, negative, memo)
	}
	if sing == positive {
		memo[i][sing] = max(nums[i]*best, nums[i])
	} else {
		memo[i][sing] = min(nums[i]*best, nums[i])
	}

	return memo[i][sing]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{-3, -1, -1}
	fmt.Print(maxProduct(nums))
}
