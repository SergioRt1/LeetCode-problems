package main

import "fmt"

func maxProfit(prices []int) int {
	var diff, best int
	min := prices[0]

	for _, price := range prices {
		diff = price - min
		if diff > best {
			best = diff
		}
		if price < min {
			min = price
		}
	}

	return best
}

func main() {
	p := []int{7}
	fmt.Print(maxProfit(p))
}
