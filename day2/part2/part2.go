package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

func main() {
	lines, solve := shared.Init(2, 2)

	horizontalPos, depth, aim := 0, 0, 0

	for _, line := range lines {
		fields := strings.Fields(line)
		i := shared.Atoi(fields[1])
		switch fields[0] {
		case "forward":
			horizontalPos += i
			depth += aim * i
		case "down":
			aim += i
		case "up":
			aim -= i
		}
	}

	solve(horizontalPos * depth)
}
