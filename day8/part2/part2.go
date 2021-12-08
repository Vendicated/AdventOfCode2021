package main

import (
	"sort"
	"strings"

	"github.com/Vendicated/AdventOfCode2021/shared"
)

var numMap = map[string]string{
	"abcefg":  "0",
	"cf":      "1",
	"acdeg":   "2",
	"acdfg":   "3",
	"bcdf":    "4",
	"abdfg":   "5",
	"abdefg":  "6",
	"acf":     "7",
	"abcdefg": "8",
	"abcdfg":  "9",
}

func main() {
	lines, solve := shared.Init(8, 2)

	result := 0
	for _, line := range lines {
		mappings := make(map[int32]int32, 7)
		fiveChars, sixChars := make([]string, 3), make([]string, 3)
		comboString, outputString := shared.Partition(line, "|")
		combos := strings.Fields(comboString)

		var one, four, seven, eight string
		{
			// First, find and store the unique length combos
			fivesCursor, sixCursor := 0, 0
			for _, combo := range combos {
				switch len(combo) {
				case 2:
					one = combo
				case 3:
					seven = combo
				case 4:
					four = combo
				case 5:
					fiveChars[fivesCursor] = combo
					fivesCursor++
				case 6:
					sixChars[sixCursor] = combo
					sixCursor++
				case 7:
					eight = combo
				default:
					panic("AAAAAAAHHHH")
				}
			}
		}

		bAndD, cAndF := "", one

		// Seven is One + 'a', so simply find the char that is in seven but not one and map it to 'a'
		for _, c := range seven {
			if !strings.Contains(one, string(c)) {
				mappings[c] = 'a'
				break
			}
		}
		// Four contains b, c, d and f. We know c and f, so the other two chars are b and d
		for _, c := range four {
			s := string(c)
			if !strings.Contains(cAndF, s) {
				bAndD += s
			}
		}
		{
			var twoAndThree []string
			// Find two and three using our info
			for _, w := range fiveChars {
				if !shared.StringContains(w, int32(bAndD[0])) || !shared.StringContains(w, int32(bAndD[1])) {
					twoAndThree = append(twoAndThree, w)
				}
			}
			// One of them has an e, the other has an f
			ef1, ef2 := getDifferingChar(twoAndThree[0], twoAndThree[1]), getDifferingChar(twoAndThree[1], twoAndThree[0])
			// One contains f so check which one of the two is in One
			if shared.StringContains(one, ef1) {
				mappings[ef1] = 'f'
				mappings[ef2] = 'e'
			} else {
				mappings[ef1] = 'e'
				mappings[ef2] = 'f'
			}
			// Now that we know f, find c with that info
			for _, c := range one {
				if mappings[c] != 'f' {
					mappings[c] = 'c'
				}
			}
		}
		{
			// We now know c and f, so the other two chars in 4 are d and b
			for _, c := range four {
				if mappings[c] != 'c' && mappings[c] != 'f' {
					// Compare the char to two different five character long numbers, since only one of the three contains a b and all contain a d
					if shared.StringContains(fiveChars[0], c) && shared.StringContains(fiveChars[1], c) {
						mappings[c] = 'd'
					} else {
						mappings[c] = 'b'
					}
				}
			}
		}
		// Now we know all characters except g, so just find the last one in eight which has every character
		{
			for _, c := range eight {
				switch mappings[c] {
				case 'a', 'b', 'c', 'd', 'e', 'f':
				default:
					mappings[c] = 'g'
					break
				}
			}
		}

		// Cool! Finished creating the character map, now just solve

		outputs := strings.Fields(outputString)
		outputStr := ""
		for _, output := range outputs {
			// These casts are hell, I hate it but sort doesn't take int32
			mapped := make([]int, len(output))
			for i, c := range output {
				mapped[i] = int(mappings[c])
			}
			sort.Ints(mapped)
			s := ""
			for _, c := range mapped {
				s += string(int32(c))
			}
			outputStr += numMap[s]
		}
		result += shared.Atoi(outputStr)
	}

	solve(result)
}

func getDifferingChar(s1, s2 string) int32 {
	for _, c := range s1 {
		if !shared.StringContains(s2, c) {
			return c
		}
	}
	return -1
}
