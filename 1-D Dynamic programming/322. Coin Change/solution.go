package main

import "fmt"

const inf = 1 << 30

func coinChange(coins []int, amount int) int {
	memo := make([]map[int]int, len(coins))
	for i := range memo {
		memo[i] = make(map[int]int)
	}
	response := recursion(coins, amount, 0, memo)
	if response == inf {
		return -1
	}
	return response
}

func recursion(coins []int, amount, current int, memo []map[int]int) int {
	if amount == 0 {
		return 0
	}
	if current >= len(coins) || amount < 0 {
		return inf
	}
	if response, exist := memo[current][amount]; exist {
		return response
	}
	memo[current][amount] = min(recursion(coins, amount, current+1, memo), recursion(coins, amount-coins[current], current, memo)+1)
	return memo[current][amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Print(coinChange([]int{1, 2, 5}, 11))
}
