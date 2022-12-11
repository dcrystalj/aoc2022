package main

import (
	"fmt"
)

func day8() {
	// lines := ReadInput("day8sample.in")
	lines := ReadInput("day8.in")
	grid := NewGrid(lines)
	vgrid := NewVisibility(grid)
	colorIt(grid, &vgrid)
	// printGrid(vgrid)
	cnt := countVisible(vgrid)
	// println(grid)

	println()
	fmt.Print(cnt)
	println()

	// grid2 := NewGrid2(lines)
	// printGrid2(grid2)
	// println()

	// colorIt2(grid, grid2)
	// println()

	// printGrid2(grid2)
	// println()

	// println(getMax(grid2))

}

func getMax(grid [][]int) int {
	nrows, ncols := len(grid), len(grid[0])
	m := 0
	for i := 1; i < nrows-1; i++ {
		for j := 1; j < ncols-1; j++ {
			if grid[i][j] > m {
				m = grid[i][j]
			}
		}
	}
	return m
}

func NewGrid(lines []string) [][]int {
	grid := make([][]int, 0, len(lines))
	for _, line := range lines {
		gline := make([]int, 0, len(line))
		for _, g := range []byte(line) {
			gline = append(gline, int(g-'0'))
		}
		grid = append(grid, gline)
	}
	return grid
}

func NewGrid2(lines []string) [][]int {
	grid := make([][]int, 0, len(lines))
	for _, line := range lines {
		gline := make([]int, 0, len(line))
		for i := 0; i < len(line); i++ {
			gline = append(gline, 1)
		}
		grid = append(grid, gline)
	}
	return grid
}

func NewVisibility(grid [][]int) [][]bool {
	visibility := make([][]bool, 0, len(grid))
	for _, line := range grid {
		vline := make([]bool, len(line))
		visibility = append(visibility, vline)
	}
	return visibility
}

func colorIt(grid [][]int, vgrid *[][]bool) {
	nrows, ncols := len(grid), len(grid[0])
	// from left
	for row := range grid {
		(*vgrid)[row][0] = true
		maxSoFar := grid[row][0]
		for col := 1; col < len(grid[row]); col++ {
			if grid[row][col] > maxSoFar {
				(*vgrid)[row][col] = true
				maxSoFar = grid[row][col]
			}
		}
	}

	// from right

	for row := range grid {
		(*vgrid)[row][ncols-1] = true
		maxSoFar := grid[row][ncols-1]
		for col := ncols - 2; col >= 0; col-- {
			if grid[row][col] > maxSoFar {
				(*vgrid)[row][col] = true
				maxSoFar = grid[row][col]
			}
		}
	}

	//from top
	for i := 0; i < ncols; i++ {
		(*vgrid)[0][i] = true
		maxSoFar := grid[0][i]
		for j := 1; j < nrows; j++ {
			if grid[j][i] > maxSoFar {
				(*vgrid)[j][i] = true
				maxSoFar = grid[j][i]
			}
		}
	}

	// //from bottom
	for i := 0; i < ncols; i++ {
		(*vgrid)[nrows-1][i] = true
		maxSoFar := grid[nrows-1][i]
		for j := nrows - 2; j >= 0; j-- {
			if grid[j][i] > maxSoFar {
				maxSoFar = grid[j][i]
				(*vgrid)[j][i] = true
			}
		}
	}
}

func colorIt2(grid [][]int, vgrid [][]int) {
	nrows, ncols := len(grid), len(grid[0])

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			// up
			l := 1
			for k := i - 1; k > 0 && grid[k][j] < grid[i][j]; k-- {
				l += 1
			}
			if l > 0 {
				vgrid[i][j] *= l
			}

			// down
			l = 1
			for k := i + 1; k < nrows-1 && grid[k][j] < grid[i][j]; k++ {
				l += 1
			}
			if l > 0 {
				vgrid[i][j] *= l
			}

			// left
			l = 1
			for k := j - 1; k > 0 && grid[i][k] < grid[i][j]; k-- {
				l += 1
			}
			if l > 0 {
				vgrid[i][j] *= l
			}
			// // right
			l = 1
			for k := j + 1; k < ncols-1 && grid[i][k] < grid[i][j]; k++ {
				l += 1
			}
			if l > 0 {
				vgrid[i][j] *= l
			}
		}
	}
}

func countVisible(vgrid [][]bool) int {
	sum := 0
	for _, row := range vgrid {
		for _, col := range row {
			if col {
				sum += 1
			}
		}
	}
	return sum
}

// func printGrid(grid [][]bool) {
// 	for _, row := range grid {
// 		println()
// 		for _, col := range row {
// 			if col {
// 				print(1)
// 			} else {
// 				print(0)
// 			}
// 		}
// 	}
// }

//	func printGrid2(grid [][]int) {
//		for _, row := range grid {
//			println()
//			for _, col := range row {
//				print(col)
//			}
//		}
//	}
func day8_2() {

}
