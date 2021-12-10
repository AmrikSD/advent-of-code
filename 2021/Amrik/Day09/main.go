package main

import (
	"fmt"
	"sort"
)

func main() {

	input := "input.txt"
	fmt.Println("part 1: ", partOne(input))
	fmt.Println("part 2: ", partTwo(input))

}
func partOne(input string) (answer int) {

	heightmap := ParseInput(input)
	_, lowpoints := FindLowPoints(heightmap)
	for _, v := range lowpoints {
		answer += v + 1
	}
	return answer
}

func partTwo(input string) int {

	heightmap := ParseInput(input)
	coords, _ := FindLowPoints(heightmap)
	visited := [][]bool{}
	for _, row := range heightmap {
		rowArr := []bool{}
		for range row {
			rowArr = append(rowArr, false)
		}
		visited = append(visited, rowArr)
	}

	areas := []int{}
	for _, c := range coords {
		areas = append(areas, findBasinArea(heightmap, visited, c))
	}

	sort.Ints(areas)

	product := 1
	for _, v := range areas[len(areas)-3:] {
		product *= v
	}
	return product

}
