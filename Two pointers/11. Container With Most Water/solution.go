package main

import "fmt"

// Use two pointer from each side and hold the biggest one, O(n) time
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	best := 0
	for i < j {
		if current := (j - i) * min(height[i], height[j]); current > best {
			best = current
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return best
}

// Checks for all possible i, j, where i < j and calculate all areas, O(n^2) time
func greedySolution(height []int) int {
	best := 0
	for i, vi := range height {
		for j := i + 1; j < len(height); j++ {
			if current := (j - i) * min(vi, height[j]); current > best {
				best = current
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 25, 1, 1, 25, 8, 3, 7}
	fmt.Print(maxArea(height))
}
