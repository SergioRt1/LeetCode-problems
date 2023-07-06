package main

import (
	"fmt"
	"sort"
)

// Use map to count occurrences and then sort by occurrences O(n log n)
func topKFrequent1(nums []int, k int) []int {
	occurrence := map[int]int{}
	for _, n := range nums {
		occurrence[n] += 1
	}
	response := make([]int, 0, len(occurrence))
	for key, _ := range occurrence {
		response = append(response, key)
	}
	sort.Slice(response, func(i, j int) bool {
		return occurrence[response[i]] > occurrence[response[j]]
	})

	return response[:k]
}

// Use map to count occurrences and then use a list to get the response O(n+k)
func topKFrequent(nums []int, k int) []int {
	occurrence := map[int]int{}
	freq := make([][]int, len(nums)+1)
	for _, n := range nums {
		occurrence[n] += 1
	}
	for n, count := range occurrence {
		freq[count] = append(freq[count], n)
	}
	response := make([]int, 0, k)

	for i := len(freq) - 1; i >= 0; i-- {
		for _, elem := range freq[i] {
			response = append(response, elem)
			if len(response) == k {
				return response
			}
		}
	}

	return response
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Print(topKFrequent(nums, k))
}
