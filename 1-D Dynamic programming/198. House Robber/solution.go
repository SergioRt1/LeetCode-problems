package main

import "fmt"

func rob(nums []int) int {
	//return recursiveHelper(nums)
	return tabulation(nums)
}

func tabulation(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	memo := make([]int, len(nums))

	//base cases
	last := len(nums) - 1
	memo[last] = nums[last]
	memo[last-1] = max(nums[last-1], nums[last])

	//decision
	for i := last - 2; i >= 0; i-- {
		memo[i] = max(memo[i+2]+nums[i], memo[i+1])
	}

	return memo[0]
}

func recursiveHelper(nums []int) int {
	memo := make([]int, 101)
	for i := range memo {
		memo[i] = -1
	}
	return dp(nums, 0, memo)
}

// the maximum amount of money you can rob in nums[idx:]
func dp(nums []int, idx int, memo []int) int {
	// base case
	if idx >= len(nums) {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	//recursive decision to rob the idx hose and move 2 or not and move 1
	memo[idx] = max(dp(nums, idx+2, memo)+nums[idx], dp(nums, idx+1, memo))
	return memo[idx]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Print(rob(nums))
}
