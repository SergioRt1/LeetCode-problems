package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	lo := 0
	m, n := len(matrix), len(matrix[0])
	hi := m*n - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		i, j := toCoord(mid, n)
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			lo = mid + 1
		} else if target < matrix[i][j] {
			hi = mid - 1
		}
	}
	return false
}

func toCoord(idx, n int) (int, int) {
	return idx / n, idx % n
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
