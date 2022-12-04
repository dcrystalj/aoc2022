package main

import "unicode"

func day3() {
	lines := ReadInput("day3.in")
	res := 0
	for _, line := range lines {
		c1, c2 := splitConpartment(line)
		cnt1 := countChars(c1)
		cnt2 := countChars(c2)

		common := commonChars(cnt1, cnt2)
		res += sumPriorities(common)
	}
	println(res)
}

func day3_2() {
	lines := ReadInput("day3.in")
	groups := []map[byte]int{}
	res := 0
	for _, line := range lines {
		cnt := countChars(line)
		groups = append(groups, cnt)
		if len(groups) == 3 {
			common1 := commonChars(groups[0], groups[1])
			common2 := commonChars(groups[2], common1)
			res += sumPriorities(common2)
			groups = []map[byte]int{}
		}
	}
	println(res)
}

func splitConpartment(line string) (string, string) {
	return line[:len(line)/2], line[len(line)/2:]
}

func countChars(conpartment string) map[byte]int {
	m := map[byte]int{}
	for i := range conpartment {
		char := conpartment[i]
		cnt, has := m[char]
		if has {
			m[char] = cnt + 1
		} else {
			m[char] = 1
		}
	}
	return m
}

func commonChars(cnt1, cnt2 map[byte]int) map[byte]int {
	common := map[byte]int{}
	for key := range cnt1 {
		_, has := cnt2[key]
		if has {
			common[key] = 1
		}
	}
	return common
}

func sumPriorities(common map[byte]int) int {
	sum := 0
	for char := range common {
		if isUpper(char) {
			sum += int(char - 'A' + 27)
		} else {
			sum += int(char - 'a' + 1)
		}
	}
	return sum
}

func isUpper(b byte) bool {
	return b == byte(unicode.ToUpper(rune(b)))
}
