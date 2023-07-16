package main

import "fmt"

// climbStairs given an n, the number of steps in a staircase it returns in how many way is possible to reach the top in steps of 1 or 2.
func climbStairs(n int) int {
	// return helperRecursion(n)
	return traversal(n)
}

func traversal(n int) int {
	if n < 3 {
		return n
	}
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2

	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func helperRecursion(n int) int {
	memo := make([]int, 46)
	for i := range memo {
		memo[i] = -1
	}
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2

	return dp(n, memo)
}

func dp(n int, memo []int) int {
	// base case
	if memo[n] != -1 {
		return memo[n]
	}
	// recursive case, decision
	memo[n] = dp(n-1, memo) + dp(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Print(climbStairs(1))
}
