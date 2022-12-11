package main

import (
	"fmt"
	"strconv"
)

type Point struct {
	row int
	col int
}

type State struct {
	head Point
	tail Point
}

func (s State) do(action byte, times int, grid [][]int) State {
	if times == 0 {
		return s
	}
	switch action {
	case 'U':
		s = s.one(Point{s.head.row + 1, s.head.col})
	case 'D':
		s = s.one(Point{s.head.row - 1, s.head.col})
	case 'L':
		s = s.one(Point{s.head.row, s.head.col - 1})
	case 'R':
		s = s.one(Point{s.head.row, s.head.col + 1})
	default:
		panic("unknown action")
	}
	grid[s.tail.row][s.tail.col] = 1
	return s.do(action, times-1, grid)
}

// moves rope for one step in the grid by pulling one side
func (s State) one(next Point) State {
	if s.head.row == next.row {
		if s.head.col < next.col { // RIGHT
			if s.tail.col < s.head.col || (s.tail.col == s.head.col && s.tail.row == s.head.row) { // 1
				return State{next, s.head}
			} else if s.tail.row > s.head.row && s.tail.col == s.head.col { // 2
				return State{next, s.tail}
			} else if s.tail.row < s.head.row && s.tail.col == s.head.col { // 3
				return State{next, s.tail}
			} else if s.tail.row > s.head.row && s.tail.col > s.head.col { // 4
				return State{next, s.tail}
			} else if s.tail.row < s.head.row && s.tail.col > s.head.col { // 5
				return State{next, s.tail}
			} else if s.tail.row == s.head.row && s.tail.col > s.head.col { // 6
				return State{next, next}
			}
		} else if s.head.col > next.col { // LEFT
			if s.tail.col > s.head.col || (s.tail.col == s.head.col && s.tail.row == s.head.row) { // 1
				return State{next, s.head}
			} else if s.tail.row > s.head.row && s.tail.col == s.head.col { // 2
				return State{next, s.tail}
			} else if s.tail.row < s.head.row && s.tail.col == s.head.col { // 3
				return State{next, s.tail}
			} else if s.tail.row > s.head.row && s.tail.col < s.head.col { // 4
				return State{next, s.tail}
			} else if s.tail.row < s.head.row && s.tail.col < s.head.col { // 5
				return State{next, s.tail}
			} else if s.tail.row == s.head.row && s.tail.col < s.head.col { // 6
				return State{next, next}
			}
		} else {
			panic("same col")
		}
	} else if s.head.col == next.col {
		if s.head.row < next.row { // DOWN
			if s.tail.row < s.head.row || (s.tail.col == s.head.col && s.tail.row == s.head.row) { // 1
				return State{next, s.head}
			} else if s.tail.row == s.head.row && s.tail.col > s.head.col { // 2
				return State{next, s.tail}
			} else if s.tail.row == s.head.row && s.tail.col < s.head.col { // 3
				return State{next, s.tail}
			} else if s.tail.row > s.head.row && s.tail.col > s.head.col { // 4
				return State{next, s.tail}
			} else if s.tail.row > s.head.row && s.tail.col < s.head.col { // 5
				return State{next, s.tail}
			} else if s.tail.row > s.head.row && s.tail.col == s.head.col { // 6
				return State{next, next}
			}
		} else if s.head.row > next.row { // UP
			if s.tail.row > s.head.row || (s.tail.col == s.head.col && s.tail.row == s.head.row) { // 1
				return State{next, s.head}
			} else if s.tail.row == s.head.row && s.tail.col > s.head.col { // 2
				return State{next, s.tail}
			} else if s.tail.row == s.head.row && s.tail.col < s.head.col { // 3
				return State{next, s.tail}
			} else if s.tail.row < s.head.row && s.tail.col > s.head.col { // 4
				return State{next, s.tail}
			} else if s.tail.row < s.head.row && s.tail.col < s.head.col { // 5
				return State{next, s.tail}
			} else if s.tail.row < s.head.row && s.tail.col == s.head.col { // 6
				return State{next, next}
			}
		} else {
			panic("same row")
		}
	} else {
		panic("not same row or col")
	}
	panic("should not be here")
}

func day9() {
	lines := ReadInput("day9.in")
	grid := NewGrid9()
	start := State{Point{5000, 5000}, Point{5000, 5000}}
	grid[start.head.row][start.head.col] = 1
	state := start
	for _, line := range lines {
		action := line[0]
		times, _ := strconv.Atoi(line[2:])
		state = state.do(action, times, grid)
	}

	fmt.Println(countVisited(grid))
}

func NewGrid9() [][]int {
	grid := make([][]int, 10000)
	for i := range grid {
		grid[i] = make([]int, 10000)
	}
	return grid
}

func countVisited(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col > 0 {
				count++
			}
		}
	}
	return count
}

func day9_2() {
	lines := ReadInput("day9.in")
	grid := NewGrid9()
	n := 10
	snake := make([]Point, 0, n)
	for i := 0; i < n; i++ {
		snake = append(snake, Point{5000, 5000})
	}
	grid[snake[0].row][snake[0].col] = 1
	for _, line := range lines {
		action := line[0]
		times, _ := strconv.Atoi(line[2:])

		for i := 0; i < times; i++ {
			snake[n-1] = do(snake[n-1], action, grid)
			for j := n - 1; j > 0; j-- {
				snake[j-1] = one(snake[j-1], snake[j], grid)
			}
			grid[snake[0].row][snake[0].col] = 1
		}
	}

	// drawGrid9(grid)
	fmt.Println(countVisited(grid))
}

func do(p Point, action byte, grid [][]int) Point {
	next := Point{0, 0}
	switch action {
	case 'U':
		next = Point{p.row + 1, p.col}
	case 'D':
		next = Point{p.row - 1, p.col}
	case 'L':
		next = Point{p.row, p.col - 1}
	case 'R':
		next = Point{p.row, p.col + 1}
	default:
		panic("unknown action")
	}
	return next
}

// moves rope for one step in the grid by pulling one side
func one(p, next Point, grid [][]int) Point {
	if absInt(p.row-next.row) <= 1 && absInt(p.col-next.col) <= 1 {
		return p
	}
	if next.row > p.row {
		p.row += 1
	}
	if next.row < p.row {
		p.row -= 1
	}
	if next.col > p.col {
		p.col += 1
	}
	if next.col < p.col {
		p.col -= 1
	}
	return p
}

func NewSmallGrid() [][]int {
	g := make([][]int, 15)
	for i := 0; i < 15; i++ {
		g[i] = make([]int, 15)
	}
	return g
}

func printSnake(s []Point, g [][]int) {
	for i := 0; i < len(s); i++ {
		g[s[i].row][s[i].col] = len(s) - i
	}
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func drawGrid9(grid [][]int) {
	for i := 5000 - 15; i < 5000+15; i++ {
		for j := 5000 - 15; j < 5000+15; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}
