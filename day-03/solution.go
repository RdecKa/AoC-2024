package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

const validNumberFormat = `[0-9]{1,3}`

var mulRe = regexp.MustCompile(
	fmt.Sprintf(`(mul\((%s),(%s)\))`, validNumberFormat, validNumberFormat))

func star1(memory string) int {
	return multiplyAndSum(*mulRe, memory)
}

func findEnabledSections(memory string) (enabledSections []string) {
	for _, s := range strings.Split(memory, "do()") {
		enabledSections = append(enabledSections, strings.Split(s, "don't()")[0])
	}
	return
}

func star2(memory string) (result int) {
	for _, s := range findEnabledSections(memory) {
		result += multiplyAndSum(*mulRe, s)
	}
	return
}

func solution(starFunc func(string) int, input int) int {
	inputFileName := util.InputFileName(input)
	memory := strings.Join(readInput(inputFileName), "")
	return starFunc(memory)
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
