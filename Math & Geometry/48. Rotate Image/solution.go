package main

import "fmt"

func mapping(iter, rotation, init, n int) (int, int) {
	switch rotation {
	case 0:
		return init + iter, n
	case 1:
		return n, n - iter
	case 2:
		return n - iter, init
	case 3:
		return init, init + iter
	}
	return 0, 0
}

func rec(matrix [][]int, init, n int) {
	if n-init < 1 {
		return
	}
	var tmp, nI, nJ int
	for i := 0; i < n-init; i++ {
		tmp = matrix[init][init+i]
		nI, nJ = init, i
		for k := 0; k < 3; k++ {
			nI, nJ = mapping(i, k, init, n)
			matrix[nI][nJ], tmp = tmp, matrix[nI][nJ]
		}
		matrix[init][init+i] = tmp
	}
	printMatrix(matrix)
	rec(matrix, init+1, n-1)
}

func rotate(matrix [][]int) {
	rec(matrix, 0, len(matrix)-1)
}

func printMatrix(m [][]int) {
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println("------------------")
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	printMatrix(matrix)
	rotate(matrix)

	printMatrix(matrix)
}
