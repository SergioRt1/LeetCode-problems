package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	i, j, idx := 0, 0, 0

	bottom, top := 0, len(matrix)
	left, right := 0, len(matrix[0])

	totalElements := top * right
	result := make([]int, totalElements)
	for idx < totalElements {
		// move from left to the right
		for ; idx < totalElements && j < right; j++ {
			result[idx] = matrix[i][j]
			idx++
		}
		bottom++
		j--
		// move from top to down
		i++
		for ; idx < totalElements && i < top; i++ {
			result[idx] = matrix[i][j]
			idx++
		}
		right--
		i--
		//move from right to left
		j--
		for ; idx < totalElements && left <= j; j-- {
			result[idx] = matrix[i][j]
			idx++
		}
		top--
		j++
		//move from down to up
		i--
		for ; idx < totalElements && bottom <= i; i-- {
			result[idx] = matrix[i][j]
			idx++
		}
		left++
		i++
		//move right
		j++
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	fmt.Print(spiralOrder(matrix))
}
