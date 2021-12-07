package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	input, solve := shared.InitNoSplit(6, 1)
	var lanternFish []int
	{
		split := strings.Split(input, ",")
		lanternFish = make([]int, len(split))
		for i, s := range split {
			lanternFish[i] = shared.Atoi(s)
		}
	}

	for counter := 0; counter < 80; counter++ {
		fishCount := len(lanternFish)
		for i := 0; i < fishCount; i++ {
			if lanternFish[i] == 0 {
				lanternFish[i] = 6
				lanternFish = append(lanternFish, 8)
			} else {
				lanternFish[i]--
			}
		}
	}

	solve(len(lanternFish))
}
