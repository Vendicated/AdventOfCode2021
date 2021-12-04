package main

import (
	"math"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(3, 1)

	arrLen := len(lines[0])
	setBitCount := make([]int, arrLen)

	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				setBitCount[i]++
			}
		}
	}

	gammaRate, epsilonRate := 0, 0

	halfLinesLen := float64(len(lines)) / 2
	for i, count := range setBitCount {
		bitVal := int(math.Pow(2, float64(arrLen-i-1)))
		if float64(count) > halfLinesLen {
			gammaRate += bitVal
		} else {
			epsilonRate += bitVal
		}
	}

	solve(gammaRate * epsilonRate)
}
