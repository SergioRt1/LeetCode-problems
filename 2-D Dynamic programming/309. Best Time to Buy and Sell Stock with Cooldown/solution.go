package main

import "fmt"

type state int

const (
	buyed state = iota
	not
	coldown
)

const unsaved = -1 << 31

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	var recursion func(s state, idx int) int

	memo := make([][]int, 3)
	for i := range memo {
		memo[i] = make([]int, len(prices))
		for j := range prices {
			memo[i][j] = unsaved
		}
	}

	recursion = func(s state, idx int) int {
		if idx >= len(prices) {
			return 0
		}
		if memo[s][idx] != unsaved {
			return memo[s][idx]
		}
		var profit int

		switch s {
		case not:
			profit = max(recursion(buyed, idx+1)-prices[idx], recursion(s, idx+1))
		case buyed:
			profit = max(recursion(coldown, idx+1)+prices[idx], recursion(s, idx+1))
		case coldown:
			profit = recursion(not, idx+1)
		}

		memo[s][idx] = profit

		return profit
	}

	return recursion(not, 0)
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}
