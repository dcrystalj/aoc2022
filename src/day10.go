package main

import (
	"strconv"
)

func day10() {
	lines := ReadInput("day10.in")

	cmds := readCommands(lines)
	i := 0
	reg := 1
	total := 0
	for cmd := range cmds {
		i++
		if i == 20 || (i-20)%40 == 0 {
			total += reg * i
		}
		reg += cmd
		if i > 220 {
			break
		}
	}
	println(total)
}

func readCommands(lines []string) chan int {
	c := make(chan int)
	go func() {
		for _, line := range lines {
			if line[:4] == "noop" {
				c <- 0
			} else if line[:4] == "addx" {
				c <- 0
				i, _ := strconv.Atoi(line[5:])
				c <- i
			} else {
				panic("Unknown command")
			}
		}
		close(c)
	}()
	return c
}

func day10_2() {
	lines := ReadInput("day10.in")
	cmds := readCommands(lines)
	i := 0
	reg := 1
	crt := []int{}
	for cmd := range cmds {
		if i%40 == 0 {
			printCRT(crt)
			crt = make([]int, 45)
		}
		burnCRT(crt, reg, i%40)
		i++
		reg += cmd
	}
	printCRT(crt)
}

func burnCRT(crt []int, reg, pos int) {
	if reg-1 <= pos && pos <= reg+1 {
		crt[pos] = 1
	}
}

func printCRT(crt []int) {
	for _, c := range crt {
		if c == 0 {
			print(".")
		} else {
			print("#")
		}
	}
	println()
}
