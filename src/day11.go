package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []uint64
	operation   func(uint64) uint64
	testDiv     uint64
	testTrue    uint64
	testFalse   uint64
	inspections uint64
}

func day11() {
	lines := ReadInput("day11.in")
	monkeys := parseMonkeys(lines)
	for i := 0; i < 20; i++ {
		play(monkeys, 3)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	fmt.Println(monkeys[0].inspections * monkeys[1].inspections)
}

func play(monkeys []monkey, worryDiv uint64) {
	for i := 0; i < len(monkeys); i++ {
		m := monkeys[i]
		for _, item := range m.items {
			worry := m.operation(item)
			worry /= worryDiv
			worry %= 9699690
			if worry%m.testDiv == 0 {
				monkeys[m.testTrue].items = append(monkeys[m.testTrue].items, worry)
			} else {
				monkeys[m.testFalse].items = append(monkeys[m.testFalse].items, worry)
			}
			monkeys[i].inspections++
		}
		monkeys[i].items = make([]uint64, 0)
	}
}

func parseMonkeys(lines []string) []monkey {
	monkeys := make([]monkey, 0)
	for i := 0; i < len(lines); i += 7 {
		m := monkey{
			items:     parseItems(lines[i+1]),
			operation: parseOp(lines[i+2]),
			testDiv:   lastInt(lines[i+3]),
			testTrue:  lastInt(lines[i+4]),
			testFalse: lastInt(lines[i+5]),
		}
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func parseItems(line string) []uint64 {
	res := make([]uint64, 0)
	items := strings.Split(line, " ")
	for _, item := range items {
		if len(item) == 0 {
			continue
		}
		if item[len(item)-1] == ',' {
			item = item[:len(item)-1]
		}
		i, err := strconv.Atoi(item)
		if err == nil {
			res = append(res, uint64(i))
		}
	}
	return res
}

func lastInt(line string) uint64 {
	items := parseItems(line)
	return items[len(items)-1]
}

func parseOp(line string) func(uint64) uint64 {
	operator := line[23:24]
	operand, err := strconv.Atoi(line[25:])
	if err == nil {
		switch operator {
		case "+":
			return func(x uint64) uint64 { return x + uint64(operand) }
		case "*":
			return func(x uint64) uint64 { return x * uint64(operand) }
		}
	} else {
		switch operator {
		case "+":
			return func(x uint64) uint64 { return x + x }
		case "*":
			return func(x uint64) uint64 { return x * x }
		}
	}
	panic("Unknown operator")
}
func day11_2() {
	lines := ReadInput("day11.in")
	monkeys := parseMonkeys(lines)
	for i := 0; i < 10000; i++ {
		play(monkeys, 1)
	}

	for _, v := range monkeys {
		fmt.Println(v.inspections)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	fmt.Println(monkeys[0].inspections * monkeys[1].inspections)
}
