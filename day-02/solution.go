package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RdecKa/AoC-2024/util"
)

type report struct {
	levels []int
}

func parseReport(line string) (report report) {
	stringValues := strings.Fields(line)
	for _, val := range stringValues {
		intVal, _ := strconv.Atoi(val)
		report.levels = append(report.levels, intVal)
	}
	return
}

func readReports(fileName string) (reports []report) {
	lines := util.FileToLines(fileName)
	for _, line := range lines {
		reports = append(reports, parseReport(line))
	}
	return
}

const minDiff = 1
const maxDiff = 3

func (r report) isGraduallyChanging(allowedMinDiff, allowedMaxDiff int) bool {
	for i := 1; i < len(r.levels); i++ {
		diff := r.levels[i] - r.levels[i-1]
		if diff < allowedMinDiff || diff > allowedMaxDiff {
			return false
		}
	}
	return true
}

func (r report) isGraduallyIncreasing() bool {
	return r.isGraduallyChanging(minDiff, maxDiff)
}

func (r report) isGraduallyDecreasing() bool {
	return r.isGraduallyChanging(-maxDiff, -minDiff)
}

func star1(reports []report) (safeReports int) {
	for _, r := range reports {
		if r.isGraduallyIncreasing() || r.isGraduallyDecreasing() {
			safeReports++
		}
	}
	return
}

func main() {
	const input = 1
	inputFileName := util.InputFileName(input)
	reports := readReports(inputFileName)

	fmt.Println(star1(reports))
}
