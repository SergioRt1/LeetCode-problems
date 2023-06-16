package main

import "fmt"

// Accumulate product in both ways, then multiply the partial accumulates in the middle
func productExceptSelf(nums []int) []int {
	size := len(nums)
	asc := make([]int, 0, size)
	desc := make([]int, size)
	response := make([]int, 0, size)

	acc := 1
	for i := 0; i < size; i++ {
		asc = append(asc, acc)
		acc *= nums[i]
	}
	acc = 1
	for i := size - 1; i >= 0; i-- {
		desc[i] = acc
		acc *= nums[i]
	}
	for i := 0; i < size; i++ {
		result := asc[i] * desc[i]
		response = append(response, result)
	}
	return response
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Print(productExceptSelf(nums))
}
