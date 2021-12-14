package main

import (
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

// Fix this mess
var (
	paths     []string
	steps     = make(map[string]*[]string)
	result    = 0
)

func main() {
	lines, solve := shared.Init(12, 1)
	for _, line := range lines {
		left, right := shared.GetLeftAndRight(line, "-")
		if val, ok := steps[left]; ok {
			*val = append(*val, right)
		} else {
			steps[left] = &[]string{right}
		}
	}

	for start, moves := range steps {
		if start == "start" {
			for _, move := range *moves {
				walkStep(move, nil)
			}
		}
	}

	println(result)
	if result != 10 {
		panic("oof")
	}

	solve("ouch")
}

func walkStep(currField string, currPath []string) {
	currPath = append(currPath, currField)
	possibleMoves, ok := steps[currField]
	if !ok {
		return
	}

loop:
	for _, move := range *possibleMoves {
		if !shared.IsUpper(move) {
			for _, path := range currPath {
				if path == move {
					continue loop
				}
			}
		}
		if move == "end" {
			paths = append(paths, strings.Join(append(currPath, move), "->"))
			result++
			continue loop
		}

		walkStep(move, currPath)
	}

	if shared.IsUpper(currField) {
		loop2:
		for start, ends := range steps {
			for _, end := range *ends {
				if end == currField && start != "start" {
					for _, path := range currPath {
						if path == start {
							continue loop2
						}
					}
					walkStep(start, append(currPath, start))
				}
			}
		}
	}
}
