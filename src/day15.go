package main

import (
	"fmt"
	"sort"
)

type Sensor struct {
	Point
}

type Beacon struct {
	Point
}
type SB struct {
	Sensor
	Beacon
}

func day15() {
	// searchRow := 10
	searchRow := 2000000
	lines := ReadInput("day15.in")
	lineRanges := make([][2]int, 0)
	sbs := make([]SB, 0)
	for _, line := range lines {
		sb := readSb(line)
		sbs = append(sbs, sb)
		rang, ok := covers(sb, searchRow)
		if ok {
			lineRanges = append(lineRanges, rang)
		}
	}

	sortRanges(lineRanges)

	cnt := countCoveredRanges(lineRanges)
	cnt -= len(beaconsInTheRow(sbs, searchRow))

	fmt.Println(cnt)
}

func beaconsInTheRow(sbs []SB, searchRow int) map[Beacon]struct{} {
	beacons := make(map[Beacon]struct{})
	for _, b := range sbs {
		if b.Beacon.row == searchRow {
			beacons[b.Beacon] = struct{}{}
		}
	}
	return beacons
}

func countCoveredRanges(ranges [][2]int) int {
	cnt := 0
	i := 0
	for i < len(ranges) {
		start, max := ranges[i][0], ranges[i][1]
		for ; i < len(ranges) && ranges[i][0] <= max; i++ {
			if ranges[i][1] > max {
				max = ranges[i][1]
			}
		}
		cnt += max - start + 1
	}
	return cnt
}

func notCovered(ranges [][2]int) []int {
	res := make([]int, 0)
	i := 0
	for i < len(ranges) {
		_, max := ranges[i][0], ranges[i][1]
		for ; i < len(ranges) && ranges[i][0] <= max; i++ {
			if ranges[i][1] > max {
				max = ranges[i][1]
			}
		}
		for j := max + 1; j < 4000000 && i < len(ranges) && j < ranges[i][0]; j++ {
			res = append(res, j)
		}
	}
	return res
}

func sortRanges(ranges [][2]int) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
}

func manhattan(sb SB) int {
	return absInt(sb.Beacon.col-sb.Sensor.col) + absInt(sb.Beacon.row-sb.Sensor.row)
}
func covers(sb SB, searchRow int) ([2]int, bool) {
	dist := manhattan(sb)
	if sb.Sensor.row > searchRow {
		dist -= sb.Sensor.row - searchRow
		if dist < 0 {
			return [2]int{0, 0}, false
		}
		return [2]int{sb.Sensor.col - dist, sb.Sensor.col + dist}, true
	} else {
		dist -= searchRow - sb.Sensor.row
		if dist < 0 {
			return [2]int{0, 0}, false
		}
		return [2]int{sb.Sensor.col - dist, sb.Sensor.col + dist}, true
	}
}

func coversMax(sb SB, searchRow int) ([2]int, bool) {
	max := 4000000
	dist := manhattan(sb)
	if sb.Sensor.row > searchRow {
		dist -= sb.Sensor.row - searchRow
		if dist < 0 {
			return [2]int{0, 0}, false
		}
		return [2]int{maxInt(sb.Sensor.col-dist, 0), minInt(sb.Sensor.col+dist, max)}, true
	} else {
		dist -= searchRow - sb.Sensor.row
		if dist < 0 {
			return [2]int{0, 0}, false
		}
		return [2]int{maxInt(sb.Sensor.col-dist, 0), minInt(sb.Sensor.col+dist, max)}, true
	}
}

func readSb(line string) SB {
	sb := SB{}
	fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sb.Sensor.col, &sb.Sensor.row, &sb.Beacon.col, &sb.Beacon.row, &sb.Beacon.col, &sb.Beacon.row)
	return sb
}

func day15_2() {
	// searchRow := 10
	lines := ReadInput("day15.in")
	sbs := make([]SB, 0)

	for _, line := range lines {
		sb := readSb(line)
		sbs = append(sbs, sb)
	}

	for i := 0; i < 4000000-1; i++ {
		lineRanges := make([][2]int, 0, len(sbs))
		for _, sb := range sbs {
			rang, ok := coversMax(sb, i)
			if ok {
				lineRanges = append(lineRanges, rang)
			}
		}
		sortRanges(lineRanges)
		c := notCovered(lineRanges)
		if len(c) > 0 {
			print(c[0]*4000000 + i)
			notCovered(lineRanges)
		}
	}
}
