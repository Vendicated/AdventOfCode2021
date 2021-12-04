package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Init(day, part int) ([]string, func(s interface{})) {
	fullDay := "day" + strconv.Itoa(day)
	fullPart := "part" + strconv.Itoa(part)
	return GetInputLines(fullDay), func(s interface{}) {
		file, err := os.OpenFile(filepath.Join(Unwraps(os.Getwd()), fullDay, fullPart+"-solution.txt"), os.O_CREATE|os.O_WRONLY, 0644)
		Check(err)
		CheckPair(fmt.Fprint(file, s))
		fmt.Println(s)
	}
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

func GetInput(day string) string {
	return strings.TrimSpace(Unwrapb(os.ReadFile(filepath.Join(Unwraps(os.Getwd()), day, "input.txt"))))
}

func GetInputLines(day string) []string {
	return strings.Split(GetInput(day), "\n")
}
