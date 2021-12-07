package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	input, solve := shared.InitNoSplit(6, 2)

	fishCountPerTimer := make([]int, 9) // Array where the element at each index is the amount of fish whose timer == index

	{
		split := strings.Split(input, ",")
		for _, s := range split {
			fishCountPerTimer[shared.Atoi(s)]++
		}
	}

	for count := 0; count < 256; count++ {
		newFish := fishCountPerTimer[0]

		for i := 0; i < 8; i++ {
			fishCountPerTimer[i] = fishCountPerTimer[i+1]
		}
		fishCountPerTimer[6] += newFish
		fishCountPerTimer[8] = newFish
	}

	result := 0
	for i := 0; i < 9; i++ {
		result += fishCountPerTimer[i]
	}

	solve(result)
}
