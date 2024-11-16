package main

import (
	"fmt"
)

func change(amount int, coins []int) int {
	return iter(amount, coins)
}

func rec(amount int, coins []int) int {
	var recursive func(i, currentAmount int) int

	memo := make([]map[int]int, len(coins))
	for i := 0; i < len(coins); i++ {
		memo[i] = make(map[int]int)
	}

	recursive = func(i, currentAmount int) int {
		if currentAmount == 0 {
			return 1
		}
		if i >= len(coins) || currentAmount < 0 {
			return 0
		}
		if v, ok := memo[i][currentAmount]; ok {
			return v
		}

		// Option 1: Include the current coin and reduce the amount
		include := recursive(i, currentAmount-coins[i])
		// Option 2: Skip the current coin and move to the next
		exclude := recursive(i+1, currentAmount)

		total := include + exclude
		memo[i][currentAmount] = total

		return total
	}

	return recursive(0, amount)
}

func iter(amount int, coins []int) int {
	memo := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		memo[i] = make([]int, amount+1)
	}
	for i := 0; i < len(coins); i++ {
		memo[i][0] = 1
	}
	for j := 1; j <= amount; j++ {
		for i := len(coins) - 1; i >= 0; i-- {
			currentAmount := j - coins[i]
			if currentAmount >= 0 {
				// Option 1: Include the current coin and reduce the amount
				memo[i][j] += memo[i][currentAmount]
			}
			if i+1 < len(coins) {
				// Option 2: Skip the current coin and move to the next
				memo[i][j] += memo[i+1][j]
			}
		}
	}
	//printMatrix(memo)
	return memo[0][amount]
}

func printMatrix(m [][]int) {
	for _, row := range m {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

func main() {
	a := 5
	c := []int{1, 2, 5}
	fmt.Println(change(a, c))

}
