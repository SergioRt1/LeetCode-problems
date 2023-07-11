package main

import "fmt"

// climbStairs given an n, the number of steps in a staircase it returns in how many way is possible to reach the top in steps of 1 or 2.
func climbStairs(n int) int {
	memo := make([]int, 46)
	for i := range memo {
		memo[i] = -1
	}
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2

	var dp func(n int) int
	dp = func(n int) int {
		// base case
		if memo[n] != -1 {
			return memo[n]
		}
		// recursive case, decision
		memo[n] = dp(n-1) + dp(n-2)
		return memo[n]
	}

	return dp(n)
}

func main() {
	fmt.Print(climbStairs(40))
}
