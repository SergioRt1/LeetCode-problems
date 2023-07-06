package main

import (
	"fmt"
)

func combinationSum2(candidates []int, target int) [][]int {
	return nil
}

// saves the number of solutions for each [candidates.length][target]
var memo = make([][]int, 0, 101)

type nothing struct{}

// Uses DP to store the number of solutions per [candidates.length][target], calculate recursively the solutions.
// Each solution is sorted using insertions sort to insert one value in a sorted list in linear time (the lookup to find where to insert is log(n)) O(1) in memory
// then the sorted solutions where hashed with FVN-1 algorithm to remove duplicates, the map of seen elements and the subList of solutions where
// allocated ones and assign to their final positions without coping them.
// --------------------------------------------------------------------------------------------------------------------
// Verdict: Timeout in [1, 1, 1, 1, 1, 1, 1, 1, 1, ...] case :(
func firstAttempt(candidates []int, target int) [][]int {
	for i := 0; i < 101; i++ {
		list := make([]int, 31)
		for j := range list {
			list[j] = -1
		}
		memo = append(memo, list)
	}
	r, _ := helper(candidates, target, 0)
	return r
}

func helper(candidates []int, target, responseSize int) ([][]int, map[int]nothing) {
	response := make([][]int, 0)
	if target == 0 { // goal reached
		return append(response, make([]int, responseSize)), make(map[int]nothing)
	}
	if len(candidates) == 0 { // empty response, no solution found
		return response, make(map[int]nothing)
	}
	seen := make(map[int]nothing)

	if memo[len(candidates)-1][target] != 0 { // There's no solutions
		// don't take current
		combinationsWithoutCurrent, subSeen := helper(candidates[1:], target, responseSize)
		for _, comb := range combinationsWithoutCurrent {
			response = append(response, comb)
		}
		seen = subSeen
	}
	if newTarget := target - candidates[0]; newTarget >= 0 && memo[len(candidates)-1][newTarget] != 0 { // take current if possible
		combinations, _ := helper(candidates[1:], newTarget, responseSize+1)
		for _, comb := range combinations {
			comb[responseSize] = candidates[0]
			currentCombination := comb[responseSize:]
			insertionSort(currentCombination)
			key := hashList(currentCombination)
			if _, ok := seen[key]; !ok {
				seen[key] = nothing{}
				response = append(response, comb)
			}
		}
	}
	memo[len(candidates)][target] = len(response)

	return response, seen
}

// hashList Use FVN-1 algorithm to generate a hash
func hashList(list []int) int {
	hash := 0
	prime := 31
	for _, n := range list {
		hash ^= n
		hash *= prime
	}
	return hash
}

// insertionSort takes a sorted list[1:] and put the first element in the right position to keep the entire list sorted
func insertionSort(list []int) {
	if len(list) <= 1 {
		return
	}
	n := list[0]
	idx := findInsertionPointBinary(list, n)
	//move elements to the left
	copy(list[:idx], list[1:idx+1])
	list[idx] = n
}

// findInsertionPointBinary in log time using binary search
func findInsertionPointBinary(list []int, n int) int {
	low, high, mid := 1, len(list)-1, 0
	for low <= high {
		mid = low + (high-low)/2
		if current := list[mid]; n == current {
			return mid - 1
		} else if n < current {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low - 1
}

// findInsertionPoint in linear time
func findInsertionPoint(list []int, n int) int {
	idx := len(list) - 1
	for i := 1; i < len(list); i++ {
		if list[i] >= n {
			idx = i - 1
			break
		}
	}
	return idx
}

func main() {
	candidates := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	target := 8
	fmt.Print(combinationSum2(candidates, target))
}
