package main

import "fmt"

type elf struct {
	start int
	end   int
}

func (e elf) contains(e2 elf) bool {
	return e.start <= e2.start && e2.end <= e.end
}

func (e elf) overlaps(e2 elf) bool {
	return (e.contains(e2) ||
		e2.contains(e) ||
		(e.start <= e2.start && e.end >= e2.start && e.end < e2.end) ||
		(e2.start <= e.start && e2.end >= e.start && e2.end < e.end))
}

func day4() {
	lines := ReadInput("day4.in")
	cnt := 0
	cntOverlaps := 0
	for _, line := range lines {
		e1 := elf{}
		e2 := elf{}
		fmt.Sscanf(line, "%d-%d,%d-%d", &e1.start, &e1.end, &e2.start, &e2.end)
		if e1.contains(e2) || e2.contains(e1) {
			cnt += 1
		}
		if e1.overlaps(e2) {
			cntOverlaps += 1
		}
	}
	println(cnt, cntOverlaps)

}

func day4_2() {

}
