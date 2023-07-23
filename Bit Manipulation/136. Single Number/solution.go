package main

import "fmt"

func singleNumber(nums []int) int {
	var response int
	for _, n := range nums {
		response ^= n
	}
	return response
}

func main() {
	nums := []int{-1, -1, -3, 4, 4, 1, 2, 1, 2}
	fmt.Print(singleNumber(nums))
}
