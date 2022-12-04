package main

func day2() {
	encryptions := []map[byte]byte{
		map[byte]byte{
			'X': 'A', 'Y': 'B', 'Z': 'C',
		},
		map[byte]byte{
			'X': 'A', 'Y': 'C', 'Z': 'B',
		},
		map[byte]byte{
			'X': 'B', 'Y': 'A', 'Z': 'C',
		},
		map[byte]byte{
			'X': 'B', 'Y': 'C', 'Z': 'A',
		},
		map[byte]byte{
			'X': 'C', 'Y': 'B', 'Z': 'A',
		},
		map[byte]byte{
			'X': 'C', 'Y': 'A', 'Z': 'B',
		},
	}
	maxScore := 0
	lines := ReadInput("day2.in")

	for _, e := range encryptions {
		s := 0
		for _, line := range lines {
			decrypted := e[line[2]]
			opponent := line[0]
			s += roundScore(opponent, decrypted)
		}
		if s > maxScore {
			maxScore = s
		}
	}
	println(maxScore)

	maxScore = 0
	for _, line := range lines {
		decrypted := decrpytPart2(line[0], line[2])
		maxScore += roundScore(line[0], decrypted)
	}
	println(maxScore)

}

func decrpytPart2(opponent, me byte) byte {
	if me == 'X' { //lose
		if opponent == 'A' {
			return 'C'
		} else if opponent == 'B' {
			return 'A'
		} else {
			return 'B'
		}
	} else if me == 'Y' {
		return opponent
	} else { // win
		if opponent == 'A' {
			return 'B'
		} else if opponent == 'B' {
			return 'C'
		} else {
			return 'A'
		}
	}
}

func roundScore(opponent, me byte) int {
	rs := 0
	if opponent == me {
		rs += 3
	}
	isWin := opponent == 'A' && me == 'B' || opponent == 'B' && me == 'C' || opponent == 'C' && me == 'A'
	if isWin {
		rs += 6
	}
	if me == 'A' {
		rs += 1
	} else if me == 'B' {
		rs += 2
	} else {
		rs += 3
	}
	return rs
}
