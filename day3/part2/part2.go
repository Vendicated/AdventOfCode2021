package main

import (
	"strconv"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(3, 2)
	bitLen := len(lines[0])

	oxygenGenRating := lines
	co2ScrubberRating := lines
	for i := 0; i < bitLen; i++ {
		oxygenGenRating = filterContenders(oxygenGenRating, i, false)
		if len(oxygenGenRating) == 1 {
			break
		}
	}
	for i := 0; i < bitLen; i++ {
		co2ScrubberRating = filterContenders(co2ScrubberRating, i, true)
		if len(co2ScrubberRating) == 1 {
			break
		}
	}

	solve(shared.Unwrapi64(strconv.ParseInt(oxygenGenRating[0], 2, 32)) * shared.Unwrapi64(strconv.ParseInt(co2ScrubberRating[0], 2, 32)))
}

func getPossiblyInvertedBit(c uint8, shouldInvert bool) uint8 {
	if shouldInvert {
		if c == '0' {
			return '1'
		}
		return '0'
	}
	return c
}

func filterContenders(lines []string, idx int, shouldInvertBit bool) []string {
	setBitCount := 0
	for _, line := range lines {
		if line[idx] == '1' {
			setBitCount++
		}
	}

	var desiredBit uint8
	if float64(setBitCount) >= float64(len(lines))/2 {
		desiredBit = getPossiblyInvertedBit('1', shouldInvertBit)
	} else {
		desiredBit = getPossiblyInvertedBit('0', shouldInvertBit)
	}

	var filteredList []string

	for _, line := range lines {
		if line[idx] == desiredBit {
			filteredList = append(filteredList, line)
		}
	}

	return filteredList
}
