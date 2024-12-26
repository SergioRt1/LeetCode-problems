package main

import (
	"fmt"
)

type Coord struct {
	X int
	Y int
}

const (
	empty = iota
	origin
	target
	visited
)

func uniquePathsIII(grid [][]int) int {
	dir := []Coord{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

	lenX := len(grid)
	lenY := len(grid[0])

	var start Coord
	numberOfEmptySlots := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == origin {
				start = Coord{i, j}
			} else if grid[i][j] == empty {
				numberOfEmptySlots += 1
			}
		}
	}
	deepth := 0
	var search func(current Coord) int
	search = func(current Coord) int {
		sum := 0
		for d := range dir {
			x, y := current.X+dir[d].X, current.Y+dir[d].Y
			if 0 <= x && x < lenX && 0 <= y && y < lenY { // Is in the matrix
				if grid[x][y] == empty { // Has not been visited
					grid[x][y] = visited
					deepth += 1
					sum += search(Coord{x, y})
					grid[x][y] = empty
					deepth -= 1
				}
				if grid[x][y] == target && deepth == numberOfEmptySlots {
					sum += 1
				}
			}
		}
		return sum
	}

	return search(start)
}

func main() {
	grid := [][]int{
		{0, 1},
		{2, 0},
	}

	fmt.Println(uniquePathsIII(grid))
}
