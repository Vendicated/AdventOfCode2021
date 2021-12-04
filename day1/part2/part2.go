package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(1, 2)

	lc := len(lines)
	var arr []int

	for i := 0; i < lc - 2; i++ {
		arr = append(arr, shared.Atoi(lines[i]) + shared.Atoi(lines[i + 1]) + shared.Atoi(lines[i + 2]))
	}

	increaseC := 0

	for i, n := range arr {
		if i > 0 && n > arr[i - 1] {
			increaseC++
		}
	}

	solve(increaseC)
}