package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/RdecKa/AoC-2024/util"
)

func readInputLists(fileName string) (listA []int, listB []int) {
	lines := util.FileToLines(fileName)
	for _, line := range lines {
		valuesInLine := strings.Fields(line)
		valueA, _ := strconv.Atoi(valuesInLine[0])
		valueB, _ := strconv.Atoi(valuesInLine[1])
		listA = append(listA, valueA)
		listB = append(listB, valueB)
	}
	return
}

func star1(listA []int, listB []int) (sumOfDiffs int) {
	sort.Ints(listA)
	sort.Ints(listB)
	for i := 0; i < len(listA); i++ {
		sumOfDiffs += util.Abs(listA[i] - listB[i])
	}
	return
}

func star2(listA []int, listB []int) (similarityScore int) {
	count := make(map[int]int)
	for _, el := range listB {
		count[el] += 1
	}
	for _, el := range listA {
		similarityScore += el * count[el]
	}
	return
}

func main() {
	inputFileName := util.InputFileName(0)
	listA, listB := readInputLists(inputFileName)
	fmt.Println(star1(listA, listB))
	fmt.Println(star2(listA, listB))
}
