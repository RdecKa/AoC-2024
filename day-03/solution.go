package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RdecKa/AoC-2024/util"
)

func readInput(fileName string) []string {
	return util.FileToLines(fileName)
}

func multiplyAndSum(re regexp.Regexp, s string) (result int) {
	matches := re.FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		a, _ := strconv.Atoi(m[2])
		b, _ := strconv.Atoi(m[3])
		result += a * b
	}
	return
}

func star1(sections []string) (result int) {
	validNumberFormat := `[0-9]{1,3}`
	mulRe := regexp.MustCompile(
		fmt.Sprintf(`(mul\((%s),(%s)\))`, validNumberFormat, validNumberFormat))
	for _, s := range sections {
		result += multiplyAndSum(*mulRe, s)
	}
	return
}

func star2(sections []string) int {
	return 0
}

func solution(starFunc func([]string) int, input int) int {
	inputFileName := util.InputFileName(input)
	sections := readInput(inputFileName)
	return starFunc(sections)
}

func Star1(input int) int {
	return solution(star1, input)
}
func Star2(input int) int {
	return solution(star2, input)
}

func main() {
	const input = 0
	fmt.Println(Star1(input))
	fmt.Println(Star2(input))
}
