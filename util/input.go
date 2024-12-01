package util

import (
	"fmt"
	"os"
	"strings"
)

// InputFileName : constructs input file name
func InputFileName(i int) string {
	suffix := ""
	if i > 0 {
		suffix = fmt.Sprintf("-%d", i)
	}
	return fmt.Sprintf("input%s.txt", suffix)
}

// FileToString : reads file into a string
func FileToString(fileName string) string {
	input, _ := os.ReadFile(fileName)
	return strings.TrimSuffix(string(input), "\n")
}

// FileToLines : reads file into a list of lines
func FileToLines(fileName string) []string {
	return strings.Split(FileToString(fileName), "\n")
}
