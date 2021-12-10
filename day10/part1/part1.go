package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

var errorScores = map[int32]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	lines, solve := shared.Init(10, 1)

	errorScore := 0

loop:
	for _, line := range lines {
		var closingBrackets []int32

		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				// All closing chars are openingChar + 2, only for '(' where ')' is +1
				if char == '(' {
					char--
				}
				closingBrackets = append(closingBrackets, char + 2)
			case ')', ']', '}', '>':
				if closingBrackets == nil {
					errorScore += errorScores[char]
					continue loop
				}
				lastIdx := len(closingBrackets) - 1
				if closingBrackets[lastIdx] != char {
					errorScore += errorScores[char]
					continue loop
				}
				closingBrackets = closingBrackets[:lastIdx]
			}
		}
	}
	solve(errorScore)
}
