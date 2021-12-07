package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type school map[int]int

func getNextDay(todaySchool school) school {

	nextSchool := make(school)

	keys := make([]int, 0, len(todaySchool))
	for k := range todaySchool {
		keys = append(keys, k)
	}

	// this needs to be sorted, then reversed as we need to iterate 7 before 0. Otherwise 6 will potentially miss some fishies
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		if k == 0 {
			nextSchool[8] = todaySchool[0]
			nextSchool[6] += todaySchool[0]
		} else {
			nextSchool[k-1] = todaySchool[k]
		}
	}
	return nextSchool
}

func parseInput(filePath string) school {

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	strArr := strings.Split(scanner.Text(), ",")

	startingSchool := make(school)
	for _, v := range strArr {
		n, _ := strconv.Atoi(v)
		startingSchool[n] = startingSchool[n] + 1
	}
	return startingSchool
}
