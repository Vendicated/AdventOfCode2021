package main

import (
	"sort"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

var completionScoresMap = map[int32]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	lines, solve := shared.Init(10, 2)

	var completionScores []int

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
				closingBrackets = append(closingBrackets, char+2)
			case ')', ']', '}', '>':
				if closingBrackets == nil {
					continue loop
				}
				lastIdx := len(closingBrackets) - 1
				if closingBrackets[lastIdx] != char {
					continue loop
				}
				closingBrackets = closingBrackets[:lastIdx]
			}
		}

		completionScore := 0
		for i := len(closingBrackets) - 1; i >= 0; i-- {
			completionScore = completionScore*5 + completionScoresMap[closingBrackets[i]]
		}
		completionScores = append(completionScores, completionScore)
	}

	sort.Ints(completionScores)

	median := completionScores[(len(completionScores)-1)/2]
	solve(median)
}
