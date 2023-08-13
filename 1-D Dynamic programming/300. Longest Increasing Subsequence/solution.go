package main

import "fmt"

const negInf = -1 << 30

func lengthOfLIS(nums []int) int {
	memo := make([]map[int]int, len(nums))

	for i := range memo {
		memo[i] = make(map[int]int)
	}

	return dp(nums, 0, negInf, memo)
}

func dp(nums []int, i, biggest int, memo []map[int]int) int {
	if i >= len(nums) {
		return 0
	}
	if response, seen := memo[i][biggest]; seen {
		return response
	}

	take := -1

	if nums[i] > biggest {
		take = dp(nums, i+1, nums[i], memo) + 1
	}

	memo[i][biggest] = max(take, dp(nums, i+1, biggest, memo))

	return memo[i][biggest]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Print(lengthOfLIS(nums))
}
