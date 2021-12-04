package main

import (
	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(1, 1)
	increaseC := 0
	for i, line := range lines {
		if i > 0 && shared.Atoi(line) > shared.Atoi(lines[i - 1]) {
			increaseC++
		}
	}

	solve(increaseC)
}