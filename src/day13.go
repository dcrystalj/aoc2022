package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Packet struct {
	items []interface{}
}

func day13() {
	lines := ReadInput("day13.in")
	res := 0
	for i := 1; i < len(lines); i += 3 {
		if areOrdered(lines[i-1], lines[i]) {
			res += i/3 + 1
			println(i/3 + 1)
			println()
		}
	}
	fmt.Println(res)
}

func areOrdered(a, b string) bool {
	println(a, b)
	p1, _ := parsePacket(a, 0, len(a))
	p2, _ := parsePacket(b, 0, len(b))
	return compare(p1, p2) > 0
	// fmt.Printf("%v", p1)
	// fmt.Printf("%v", p2)
	// println()
	// println()
}

func compare(p1, p2 interface{}) int {
	switch p1.(type) {
	case int:
		switch p2.(type) {
		case int:
			if p1.(int) > p2.(int) {
				return -1
			} else if p1.(int) < p2.(int) {
				return 1
			} else {
				return 0
			}
		case Packet:
			return compare(Packet{[]interface{}{p1}}, p2)
		}
	case Packet:
		switch p2.(type) {
		case int:
			return compare(p1, Packet{[]interface{}{p2}})
		case Packet:
			i, j := 0, 0
			for ; i < len(p1.(Packet).items) && j < len(p2.(Packet).items); i, j = i+1, j+1 {
				c := compare(p1.(Packet).items[i], p2.(Packet).items[j])
				if c != 0 {
					return c
				}
			}
			if len(p1.(Packet).items) < len(p2.(Packet).items) {
				return 1
			} else if len(p1.(Packet).items) > len(p2.(Packet).items) {
				return -1
			} else {
				return 0
			}
		}
	}
	return 0
}

func parsePacket(line string, start, end int) (interface{}, int) {
	p := Packet{}
	p.items = make([]interface{}, 0)
	i := start
	for ; i < end; i++ {
		if line[i] == '[' {
			pp, j := parsePacket(line, i+1, end)
			p.items = append(p.items, pp)
			i = j
		} else if line[i] == ']' {
			break
		} else if line[i] == ',' {
			continue
		} else { // number
			j := i
			for ; j < len(line); j++ {
				if line[j] == ',' || line[j] == ']' {
					break
				}
			}
			val, err := strconv.Atoi(line[i:j])
			if err != nil {
				panic(line[i:j])
			}
			p.items = append(p.items, val)
			i = j - 1
		}
	}
	return p, i + 1
}

func day13_2() {
	lines := ReadInput("day13.in")
	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")

	divider1, _ := parsePacket(lines[len(lines)-2], 0, 5)
	divider2, _ := parsePacket(lines[len(lines)-1], 0, 5)

	packets := make([]interface{}, 0)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		parsed, _ := parsePacket(lines[i], 0, len(lines[i]))
		packets = append(packets, parsed)
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) > 0
	})

	res := 1
	for i := 0; i < len(packets); i++ {
		// fmt.Printf("%v \n", packets[i])
		if compare(divider1, packets[i]) == 0 {
			println(i)
			res *= i + 1
		}
		if compare(divider2, packets[i]) == 0 {
			println(i)

			res *= i + 1
		}
	}
	fmt.Println(res)
}
