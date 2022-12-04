package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput(day string) []string {
	contents, err := ioutil.ReadFile("../data/" + day)
	Check(err)
	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func day1() {
	lines := ReadInput("day1.in")
	maxCum := 0
	curCum := 0
	for _, l := range lines {
		if l == "" {
			if curCum > maxCum {
				maxCum = curCum
			}
			curCum = 0
			continue
		}
		value, _ := strconv.Atoi(l)
		curCum += value
	}
	println(maxCum)
}

func day1_2() {
	lines := ReadInput("day1.in")
	cums := make([]int, 0, len(lines))
	curCum := 0
	for _, l := range lines {
		if l == "" {
			cums = append(cums, curCum)
			curCum = 0
			continue
		}
		value, _ := strconv.Atoi(l)
		curCum += value
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cums)))
	println(cums[0] + cums[1] + cums[2])
}
