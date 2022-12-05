package main

import "fmt"

//     [C]             [L]         [T]
//     [V] [R] [M]     [T]         [B]
//     [F] [G] [H] [Q] [Q]         [H]
//     [W] [L] [P] [V] [M] [V]     [F]
//     [P] [C] [W] [S] [Z] [B] [S] [P]
// [G] [R] [M] [B] [F] [J] [S] [Z] [D]
// [J] [L] [P] [F] [C] [H] [F] [J] [C]
// [Z] [Q] [F] [L] [G] [W] [H] [F] [M]
//  1   2   3   4   5   6   7   8   9

func day5() {
	stack := [10][]byte{
		[]byte("ZJG"),
		[]byte("QLRPWFVC"),
		[]byte("FPMCLGR"),
		[]byte("LFBWPHM"),
		[]byte("GCFSVQ"),
		[]byte("WHJZMQTL"),
		[]byte("HFSBV"),
		[]byte("FJZS"),
		[]byte("MCDPFHBT"),
	}
	lines := ReadInput("day5.in")
	instructions := parseInputs(lines)
	for _, instruction := range instructions {
		move(&stack, instruction[1], instruction[2], instruction[0])
	}

	for i := 0; i < 10; i++ {
		if len(stack[i]) > 0 {
			fmt.Printf("%c", stack[i][len(stack[i])-1])
		}
	}
	println()
}

func parseInputs(lines []string) [][3]int {
	instructions := make([][3]int, 0, len(lines))
	for _, line := range lines {
		instruction := [3]int{}
		fmt.Sscanf(line, "move %d from %d to %d", &instruction[0], &instruction[1], &instruction[2])
		instructions = append(instructions, instruction)
	}
	return instructions
}

func move(s *[10][]byte, from, to, amount int) {
	from -= 1
	to -= 1
	for i := 0; i < amount; i++ {
		toMove := s[from][len(s[from])-1]
		s[to] = append(s[to], toMove)
		s[from] = s[from][:len(s[from])-1]
	}
}

func day5_2() {
	stack := [10][]byte{
		[]byte("ZJG"),
		[]byte("QLRPWFVC"),
		[]byte("FPMCLGR"),
		[]byte("LFBWPHM"),
		[]byte("GCFSVQ"),
		[]byte("WHJZMQTL"),
		[]byte("HFSBV"),
		[]byte("FJZS"),
		[]byte("MCDPFHBT"),
	}
	lines := ReadInput("day5.in")
	instructions := parseInputs(lines)
	for _, instruction := range instructions {
		move2(&stack, instruction[1], instruction[2], instruction[0])
	}

	for i := 0; i < 10; i++ {
		if len(stack[i]) > 0 {
			fmt.Printf("%c", stack[i][len(stack[i])-1])
		}
	}
}

func move2(s *[10][]byte, from, to, amount int) {
	from -= 1
	to -= 1
	s[to] = append(s[to], s[from][len(s[from])-amount:]...)
	s[from] = s[from][:len(s[from])-amount]
}
