package main

import "fmt"

func subsets(nums []int) [][]int {
	return rec(nums, 0)
}

// Recursive solution with arrays
func rec(nums []int, idx int) [][]int {
	response := make([][]int, 0)
	if idx == len(nums) {
		return append(response, []int{})
	}

	for _, elem := range rec(nums, idx+1) {
		include := make([]int, len(elem))
		copy(include, elem)
		include = append(include, nums[idx])
		response = append(response, elem, include)
	}

	return response
}

func rec2(nums []int) [][]int {
	response := make([][]int, 0)
	if len(nums) == 0 {
		return append(response, []int{})
	}

	for _, elem := range rec2(nums[1:]) {
		include := make([]int, len(elem))
		copy(include, elem)
		include = append(include, nums[0])
		response = append(response, elem, include)
	}

	return response
}

// Use bit set to decode groups
func bits(nums []int, idx int) []int {
	if idx == len(nums) {
		return []int{0}
	}
	response := make([]int, 0)
	for _, elem := range bits(nums, idx+1) {
		include := 1<<(nums[idx]+10) | elem
		response = append(response, elem, include)
	}

	return response
}

func decode(nums []int) [][]int {
	response := make([][]int, 0, len(nums))
	for _, n := range nums {
		base := -10
		set := make([]int, 0)
		for n > 0 {
			if n&1 == 1 {
				set = append(set, base)
			}
			base++
			n >>= 1
		}
		response = append(response, set)
	}
	return response
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
	fmt.Println(bits(nums, 0))
	fmt.Println(decode(bits(nums, 0)))
}
