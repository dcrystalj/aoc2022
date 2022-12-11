package main

func day6() {
	lines := ReadInput("day6.in")

	println(running_unique_4([]byte(lines[0])))
	println(running_unique_14([]byte(lines[0])))

}

func running_unique_4(s []byte) int {
	running := map[byte]int{}
	for _, c := range s[:4] {
		running[c] += 1
	}
	if len(running) == 4 {
		return 4
	}

	for i, c := range s[4:] {
		running[byte(c)] += 1
		running[s[i]] -= 1
		if running[s[i]] == 0 {
			delete(running, s[i])
		}
		if len(running) == 4 {
			return i + 5
		}
	}
	return -1
}

func running_unique_14(s []byte) int {
	running := map[byte]int{}
	for _, c := range s[:14] {
		running[c] += 1
	}
	if len(running) == 14 {
		return 14
	}

	for i, c := range s[14:] {
		running[byte(c)] += 1
		running[s[i]] -= 1
		if running[s[i]] == 0 {
			delete(running, s[i])
		}
		if len(running) == 14 {
			return i + 15
		}
	}
	return -1
}

func day6_2() {

}
