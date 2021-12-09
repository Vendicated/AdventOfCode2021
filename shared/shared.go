package shared

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func makeSaveFunc(day, part string) func(s interface{}) {
	return func(s interface{}) {
		file, err := os.OpenFile(filepath.Join(Unwraps(os.Getwd()), day, part+"-solution.txt"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		Check(err)
		CheckPair(fmt.Fprint(file, s))
		fmt.Println(s)
	}
}

func Init(day, part int) ([]string, func(s interface{})) {
	fullDay := "day" + strconv.Itoa(day)
	fullPart := "part" + strconv.Itoa(part)
	return GetInputLines(fullDay), makeSaveFunc(fullDay, fullPart)
}

func InitNoSplit(day, part int) (string, func(s interface{})) {
	fullDay := "day" + strconv.Itoa(day)
	fullPart := "part" + strconv.Itoa(part)
	return GetInput(fullDay), makeSaveFunc(fullDay, fullPart)
}

func CheckPair(ret interface{}, e error) {
	Check(e)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Assert(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func Unwraps(s string, e error) string {
	Check(e)
	return s
}

func Unwrapb(b []byte, e error) string {
	return Unwraps(string(b), e)
}

func Unwrapi(i int, e error) int {
	Check(e)
	return i
}

func Unwrapi64(i int64, e error) int64 {
	Check(e)
	return i
}

func Atoi(str string) int {
	return Unwrapi(strconv.Atoi(str))
}

func StringContains(str string, char int32) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}

func Partition(s, separator string) (string, string) {
	split := strings.Split(s, separator)
	return split[0], split[1]
}

func GetInput(day string) string {
	return strings.TrimSpace(Unwrapb(os.ReadFile(filepath.Join(Unwraps(os.Getwd()), day, "input.txt"))))
}

func GetInputLines(day string) []string {
	return strings.Split(GetInput(day), "\n")
}

func GetAs2dArray(lines []string) [][]int {
	arr := make([][]int, len(lines))
	for i, line := range lines {
		arr[i] = GetNumbers(line, "")
	}
	return arr
}

func GetNumbers(s, separator string) []int {
	split := strings.Split(s, separator)
	nums := make([]int, len(split))
	for i, num := range split {
		nums[i] = Atoi(num)
	}
	return nums
}

func GetMinAndMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max := math.MaxInt, math.MinInt
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}