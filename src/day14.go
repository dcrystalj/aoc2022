package main

import (
	"strconv"
	"strings"
)

type Line struct {
	p1, p2 Point
}

func day14() {
	lines := ReadInput("day14.in")
	g := NewGrid14()
	for _, line := range lines {
		lines := readLine(line)
		drawLines(g, lines)
	}
	// printGrid14(g)

	cnt := 0
	for i := 0; i < 10000; i++ {
		// println(i)
		// printGrid14(g)

		if pour(g) {
			cnt++
		} else {
			break
		}
	}
	println(cnt)
	// printGrid14(g)

}

func printGrid14(g [][]int) {
	for i, row := range g {
		if i > 1000 {
			break
		}
		for j, col := range row {
			if j > 1000 {
				break
			}
			print(col)
		}
		println()
	}
}

func pour(g [][]int) bool {
	pourCol := 500
	row := 0
	if g[row][pourCol] > 0 {
		return false
	}
	// down as much as we can
	for row < len(g)-2 {
		if g[row+1][pourCol] == 0 {
			row += 1
			continue
		}
		if g[row+1][pourCol-1] == 0 {
			row += 1
			pourCol -= 1
			continue
		}
		if g[row+1][pourCol+1] == 0 {
			row += 1
			pourCol += 1
			continue
		}
		g[row][pourCol] = 2
		return true
	}
	return false
}

func NewGrid14() [][]int {
	grid := make([][]int, 10000)
	for i := 0; i < 10000; i++ {
		grid[i] = make([]int, 10000)
	}
	return grid
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func drawLine(g [][]int, line Line) int {
	if line.p1.col == line.p2.col {
		for row := minInt(line.p1.row, line.p2.row); row <= maxInt(line.p2.row, line.p1.row); row++ {
			g[row][line.p1.col] = 1
		}
	} else if line.p1.row == line.p2.row {
		for col := minInt(line.p1.col, line.p2.col); col <= maxInt(line.p2.col, line.p1.col); col++ {
			g[line.p1.row][col] = 1
		}
	} else {
		panic("error")
	}
	return maxInt(line.p1.row, line.p2.row)
}

func drawLines(g [][]int, lines []Line) int {
	max_row := 0
	for _, line := range lines {
		mrow := drawLine(g, line)
		if mrow > max_row {
			max_row = mrow
		}
	}
	return max_row
}

func readLine(line string) []Line {
	var lines []Line
	raw_points := strings.Split(line, " -> ")
	points := []Point{}
	for _, raw_point := range raw_points {
		points = append(points, readPoint(raw_point))
	}

	for i := 1; i < len(points); i++ {
		lines = append(lines, Line{points[i-1], points[i]})
	}
	return lines
}

func readPoint(raw_point string) Point {
	parts := strings.Split(raw_point, ",")
	part1, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	part2, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return Point{part2, part1 - 0}
}

func day14_2() {
	lines := ReadInput("day14.in")
	g := NewGrid14()
	maxRow := 0
	for _, line := range lines {
		lines := readLine(line)
		mrow := drawLines(g, lines)
		if mrow > maxRow {
			maxRow = mrow
		}
	}
	maxRow += 2
	// printGrid14(g)
	drawLine(g, Line{Point{maxRow, 0}, Point{maxRow, 9000}})

	cnt := 0
	for i := 0; i < 1000000; i++ {
		// println(i)
		// printGrid14(g)

		if pour(g) {
			cnt++
		} else {
			break
		}
	}
	printGrid14(g)
	println(cnt)
}
