package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		set[n] = struct{}{}
	}
	best, length := 0, 0
	for _, n := range nums {
		if _, hasBefore := set[n-1]; !hasBefore {
			length = 1
			for {
				if _, hasNext := set[n+1]; !hasNext {
					break
				}
				length++
				n++
			}
			if length > best {
				best = length
			}
		}
	}

	return best
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Print(longestConsecutive(nums))
}
