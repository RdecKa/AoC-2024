package main

import (
	"fmt"
	"strings"

	"github.com/RdecKa/AoC-2024/util"
	"github.com/dlclark/regexp2"
)

const xmasLen = 4

func constructRegexs(gridSize int) [8]*regexp2.Regexp {
	reRight := regexp2.MustCompile(`(XMAS)`, 0)
	reLeft := regexp2.MustCompile(`(SAMX)`, 0)

	vertBreak := fmt.Sprintf(`[\S\s]{%d}`, gridSize)
	reDown := regexp2.MustCompile(fmt.Sprintf(`(?=(X%sM%sA%sS))`, vertBreak, vertBreak, vertBreak), 0)
	reUp := regexp2.MustCompile(fmt.Sprintf(`(?=(S%sA%sM%sX))`, vertBreak, vertBreak, vertBreak), 0)

	// Diagonal top-left <=> botton-right
	diagRBreak := fmt.Sprintf(`[\S\s]{%d}`, gridSize+1)
	ensureRSpace := fmt.Sprintf(`[\S]{%d}[\S\s]{%d}`, xmasLen-1, gridSize+1-(xmasLen-1))
	reRightDown := regexp2.MustCompile(fmt.Sprintf(`(?=(X%sM%sA%sS))`, ensureRSpace, diagRBreak, diagRBreak), 0)
	reLeftUp := regexp2.MustCompile(fmt.Sprintf(`(?=(S%sA%sM%sX))`, ensureRSpace, diagRBreak, diagRBreak), 0)

	// Diagonal top-rigt <=> botton-left
	diagLBreak := fmt.Sprintf(`[\S\s]{%d}`, gridSize-1)
	ensureLSpace := fmt.Sprintf(`[\S\s]{%d}`, xmasLen-1)
	reLeftDown := regexp2.MustCompile(fmt.Sprintf(`(?=(%sX%sM%sA%sS))`, ensureLSpace, diagLBreak, diagLBreak, diagLBreak), 0)
	reRightUp := regexp2.MustCompile(fmt.Sprintf(`(?=(%sS%sA%sM%sX))`, ensureLSpace, diagLBreak, diagLBreak, diagLBreak), 0)

	result := [8]*regexp2.Regexp{reRight, reLeft, reDown, reUp, reRightDown, reLeftUp, reLeftDown, reRightUp}
	return result
}

func countMatches(re *regexp2.Regexp, s string) (count int) {
	m, _ := re.FindStringMatch(s)
	for m != nil {
		count++
		m, _ = re.FindNextMatch(m)
	}
	return
}

func star1(inputString string, gridSize int) (count int) {
	res := constructRegexs(gridSize)
	for _, re := range res {
		count += countMatches(re, inputString)
	}
	return
}

func constructRegexs2(gridSize int) [4]*regexp2.Regexp {
	re1 := regexp2.MustCompile(fmt.Sprintf(`(?=(M[\S]M[\S\s]{%d}A[\S\s]{%d}S[\S]S))`, gridSize-1, gridSize-1), 0)
	re2 := regexp2.MustCompile(fmt.Sprintf(`(?=(M[\S]S[\S\s]{%d}A[\S\s]{%d}M[\S]S))`, gridSize-1, gridSize-1), 0)
	re3 := regexp2.MustCompile(fmt.Sprintf(`(?=(S[\S]M[\S\s]{%d}A[\S\s]{%d}S[\S]M))`, gridSize-1, gridSize-1), 0)
	re4 := regexp2.MustCompile(fmt.Sprintf(`(?=(S[\S]S[\S\s]{%d}A[\S\s]{%d}M[\S]M))`, gridSize-1, gridSize-1), 0)
	result := [4]*regexp2.Regexp{re1, re2, re3, re4}
	return result
}

func star2(inputString string, gridSize int) (count int) {
	res := constructRegexs2(gridSize)
	for _, re := range res {
		count += countMatches(re, inputString)
	}
	return
}

func gridSize(inputString string) int {
	return len(strings.Split(inputString, "\n")[0])
}

func solution(starFunc func(string, int) int, input int) int {
	inputFileName := util.InputFileName(input)
	inputString := util.FileToString(inputFileName)
	return starFunc(inputString, gridSize(inputString))
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
