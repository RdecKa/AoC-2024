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

func (r report) excludeLevel(idx int) (newReport report) {
	for i := range r.levels {
		if i == idx {
			continue
		}
		newReport.levels = append(newReport.levels, r.levels[i])
	}
	return
}

const minDiff = 1
const maxDiff = 3

func (r report) isGraduallyChanging(allowedMinDiff, allowedMaxDiff int, allowJoker bool) bool {
	for i := 1; i < len(r.levels); i++ {
		diff := r.levels[i] - r.levels[i-1]
		if diff < allowedMinDiff || diff > allowedMaxDiff {
			if !allowJoker {
				return false
			}
			return r.excludeLevel(i).isGraduallyChanging(allowedMinDiff, allowedMaxDiff, false) ||
				r.excludeLevel(i-1).isGraduallyChanging(allowedMinDiff, allowedMaxDiff, false)
		}
	}
	return true
}

func (r report) isGraduallyIncreasing(allowJoker bool) bool {
	return r.isGraduallyChanging(minDiff, maxDiff, allowJoker)
}

func (r report) isGraduallyDecreasing(allowJoker bool) bool {
	return r.isGraduallyChanging(-maxDiff, -minDiff, allowJoker)
}

func star1(reports []report) (safeReports int) {
	for _, r := range reports {
		if r.isGraduallyIncreasing(false) || r.isGraduallyDecreasing(false) {
			safeReports++
		}
	}
	return
}

func star2(reports []report) (safeReports int) {
	for _, r := range reports {
		if r.isGraduallyIncreasing(true) || r.isGraduallyDecreasing(true) {
			safeReports++
		}
	}
	return
}

func main() {
	const input = 0
	inputFileName := util.InputFileName(input)
	reports := readReports(inputFileName)

	fmt.Println(star1(reports))
	fmt.Println(star2(reports))
}
