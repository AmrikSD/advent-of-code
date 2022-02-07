package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func getDistanceForCrabsToHorizonalSpaceP1(crabs map[int]int) int {

	distanceForCrabs := make(map[int]int)

	minPos := math.MaxInt
	maxPos := math.MinInt

	for k := range crabs {
		if k > maxPos {
			maxPos = k
		}
		if k < minPos {
			minPos = k
		}
	}
	for pos := minPos; pos <= maxPos; pos++ {
		distanceForCrabs[pos] = 0
		for k, v := range crabs {
			dist := (pos - k)
			if dist < 0 {
				dist = dist * -1
			}
			distanceForCrabs[pos] += dist * v
		}

	}
	min := math.MaxInt

	for _, v := range distanceForCrabs {
		if v < min {
			min = v
		}
	}

	return min

}

func getDistanceForCrabsToHorizonalSpaceP2(crabs map[int]int) int {

	minPos := math.MaxInt
	maxPos := math.MinInt

	for k := range crabs {
		if k > maxPos {
			maxPos = k
		}
		if k < minPos {
			minPos = k
		}
	}

	min := math.MaxInt
	for pos := minPos; pos < maxPos; pos++ {
		posFuelUsed := 0
		for k, v := range crabs {
			dist := (pos - k)
			if dist < 0 {
				dist = dist * -1
			}
			posFuelUsed += (dist * (dist + 1)) / 2 * v
		}
		if posFuelUsed < min {
			min = posFuelUsed
		}

	}

	return min

}

func parseInput(filePath string) map[int]int {

	crabs := make(map[int]int)

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()

	input := strings.Split(text, ",")

	for _, v := range input {
		num, _ := strconv.Atoi(v)
		crabs[num]++
	}

	return crabs
}
