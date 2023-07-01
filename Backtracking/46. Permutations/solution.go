package main

import "fmt"

func permute(nums []int) [][]int {
	response := make([][]int, 0)
	// base case, return
	if len(nums) == 0 {
		return append(response, []int{})
	}

	// recursive decision
	// where to put the current value (check all positions)
	current := nums[0]
	for _, subPermutation := range permute(nums[1:]) {
		for i := range nums {
			permutation := make([]int, 0, len(nums))
			used := 0
			for j := 0; j+used < len(nums); {
				if i == j && used == 0 {
					permutation = append(permutation, current)
					used++
				} else {
					permutation = append(permutation, subPermutation[j])
					j++
				}
			}
			response = append(response, permutation)
		}
	}
	return response
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Print(permute(nums))
}
