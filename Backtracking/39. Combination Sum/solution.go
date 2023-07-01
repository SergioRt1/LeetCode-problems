package main

import "fmt"

// combinationSum take a list of candidates, a target to find de combinationSum, it returns the all combinations
func combinationSum(candidates []int, target int) [][]int {
	response := make([][]int, 0)
	if target == 0 {
		return append(response, []int{})
	}
	if len(candidates) == 0 {
		return response
	}
	// decision:
	// use current element
	if newTarget := target - candidates[0]; newTarget >= 0 {
		for _, elem := range combinationSum(candidates, newTarget) {
			elem = append(elem, candidates[0])
			response = append(response, elem)
		}
	}
	// Skip current element
	for _, elem := range combinationSum(candidates[1:], target) {
		response = append(response, elem)
	}

	return response
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Print(combinationSum(candidates, target))
}
