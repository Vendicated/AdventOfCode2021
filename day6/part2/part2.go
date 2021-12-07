package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	input, solve := shared.InitNoSplit(6, 2)

	fishMap := map[int]int{ // obsolete and same as make(map[int]int, 9), but this makes it clearer what is being done here.
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	{
		split := strings.Split(input, ",")
		for _, s := range split {
			fishMap[shared.Atoi(s)]++
		}
	}

	for count := 0; count < 256; count++ {
		newFish := fishMap[0]

		for i := 0; i < 8; i++ {
			fishMap[i] = fishMap[i + 1]
		}
		fishMap[6] += newFish
		fishMap[8] = newFish
	}

	result := 0
	for i := 0; i < 9; i++ {
		result += fishMap[i]
	}

	solve(result)
}
