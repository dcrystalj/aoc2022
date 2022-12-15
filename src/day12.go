package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type DPoinHeap []*DPoint

func (h DPoinHeap) Len() int           { return len(h) }
func (h DPoinHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (pq DPoinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *DPoinHeap) Push(x any) {
	n := len(*pq)
	item := x.(*DPoint)
	item.index = n
	*pq = append(*pq, item)
}

func (h *DPoinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *DPoinHeap) update(item *DPoint, priority int) {
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

// type queue[T any] struct {
// 	bucket []T
// }

// func newQueue[T any]() *queue[T] {
// 	return &queue[T]{
// 		bucket: []T{},
// 	}
// }

// func (q *queue[T]) append(input T) {
// 	q.bucket = append(q.bucket, input)
// }

// func (q *queue[T]) tryDequeue() (T, bool) {
// 	if len(q.bucket) == 0 {
// 		var dummy T
// 		return dummy, false
// 	}
// 	value := q.bucket[0]
// 	var zero T
// 	q.bucket[0] = zero // Avoid memory leak
// 	q.bucket = q.bucket[1:]
// 	return value, true
// }

// dijkstra's algorithm from wikipedia
// procedure uniform_cost_search(start) is
//     node ← start
//     pq ← priority queue containing node only
//     set ← empty set
//     do
//         if pq is empty then
//             return failure
//         node ← pq.pop()
//         if node is a goal state then
//             return solution(node)
//         set.add(node)
//         for each of node's neighbors n do
//             if n is not in set and not in pq then
//                 pq.add(n)
//             else if n is in pq with higher cost
//                 replace existing node with n

var moves = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type DPoint struct {
	row, col int
	priority int
	prev     *DPoint
	index    int
}

func day12() {
	lines := ReadInput("day12.in")
	start, end := getStartAndEndPoint(lines)
	res := dijkstra(start, end, lines)
	fmt.Println(res)
}

func dijkstra(start Point, end Point, lines []string) int {
	ncols := len(lines[0])
	q := make(DPoinHeap, 0)
	q.Push(&DPoint{start.row, start.col, 0, nil, 0})

	heap.Init(&q)

	visited := make([]bool, ncols*len(lines))
	for {
		if len(q) == 0 {
			return 10000
		}
		p, ok := heap.Pop(&q).(*DPoint)
		if !ok {
			panic("ERRO")
		}
		if visited[p.row*ncols+p.col] {
			continue
		}
		visited[p.row*ncols+p.col] = true
		for _, m := range moves {
			np := Point{col: p.col + m.col, row: p.row + m.row}
			if np.row == end.row && np.col == end.col && lines[p.row][p.col] == 'z' {

				dist := p.priority
				// g := createGrid(lines)
				// for p.prev != nil {
				// 	fmt.Printf("%#v \n", p)
				// 	g[p.row][p.col] = '*'
				// 	p = p.prev
				// }
				// printGrid(g)
				return dist + 1
			}
			if np.col < 0 || np.row >= len(lines) || np.row < 0 || np.col >= len(lines[0]) {
				continue
			}
			if !visited[np.row*ncols+np.col] && ((lines[p.row][p.col] == 'S' && lines[np.row][np.col] == 'a') || (int(lines[np.row][np.col]) <= (int(lines[p.row][p.col]) + 1))) {
				heap.Push(&q, &DPoint{np.row, np.col, p.priority + 1, p, 0})
			}
		}
	}
	return 0
}

func createGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for row, line := range lines {
		grid[row] = make([]byte, len(line))
		for col, c := range line {
			grid[row][col] = byte(c)
		}
	}
	return grid
}

func printGrid(grid [][]byte) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func getStartAndEndPoint(lines []string) (Point, Point) {
	start, end := Point{}, Point{}
	for row, line := range lines {
		i := strings.Index(line, "S")
		if i >= 0 {
			start = Point{row, i}
		}
		i = strings.Index(line, "E")
		if i >= 0 {
			end = Point{row, i}
		}
	}
	return start, end
}

func getStartPoints(lines []string) chan Point {
	ch := make(chan Point)
	go func() {
		for row, line := range lines {
			for col, c := range line {
				if c == 'a' {
					ch <- Point{row, col}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func day12_2() {
	lines := ReadInput("day12.in")
	min := 100000
	_, end := getStartAndEndPoint(lines)
	for start := range getStartPoints(lines) {
		res := dijkstra(start, end, lines)
		if res < min {
			min = res
		}
	}
	fmt.Println(min)
}
