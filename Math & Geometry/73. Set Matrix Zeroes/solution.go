package main

import "fmt"

func setZeroes(matrix [][]int) {
	foundZero := false
	var firstRow, firstCol int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				if !foundZero {
					foundZero = true
					firstRow = i
					firstCol = j
				}
				matrix[firstRow][j] = 0
				matrix[i][firstCol] = 0
			}
		}
	}
	if foundZero {
		for i := range matrix {
			if i == firstRow {
				continue
			}
			for j := range matrix[i] {
				if j == firstCol {
					continue
				}
				if matrix[firstRow][j] == 0 || matrix[i][firstCol] == 0 {
					matrix[i][j] = 0
				}
			}
		}
		for i := range matrix {
			matrix[i][firstCol] = 0
		}
		for j := range matrix[firstRow] {
			matrix[firstRow][j] = 0
		}
	}
}

func printMatrix(m [][]int) {
	for idx := range m {
		fmt.Println(m[idx])
	}
	fmt.Println("-----------------------------")
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	printMatrix(matrix)
	setZeroes(matrix)
	printMatrix(matrix)
}
