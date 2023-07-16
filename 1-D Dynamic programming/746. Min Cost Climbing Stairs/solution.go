package main

import "fmt"

const Inf = 1 << 30

// minCostClimbingStairs calculate the minimum cost to reach the top of the floor
func minCostClimbingStairs(cost []int) int {
	return tabulation(cost)
}

func tabulation(cost []int) int {
	memo := make([]int, len(cost))
	last := len(cost) - 1
	memo[last] = cost[last]
	memo[last-1] = cost[last-1]
	for i := last - 2; i >= 0; i-- {
		memo[i] = min(memo[i+1], memo[i+2]) + cost[i]
	}
	return min(memo[0], memo[1])
}

func recursionHelper(cost []int) int {
	memo := make([]int, 1001)
	for i := range memo {
		memo[i] = -1
	}

	return min(dp(cost, 0, memo), dp(cost, 1, memo))
}

func dp(cost []int, idx int, memo []int) int {
	// base case
	if idx >= len(cost) { // Goal reached
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}

	memo[idx] = min(dp(cost, idx+1, memo), dp(cost, idx+2, memo)) + cost[idx]

	return memo[idx]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Print(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
